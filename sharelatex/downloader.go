package sharelatex

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/go-resty/resty"
)

const (
	loginURL            = "https://www.sharelatex.com/login"
	projectZipURLFormat = "https://www.sharelatex.com/project/%s/download/zip"
	sessionKey          = "sharelatex_session"
)

// Downloader interface for downloading sharelatex project
type Downloader interface {
	Download(wd string) error
}

// NewDownloader create a new downloader instance
func NewDownloader(email, password, projectURL string) Downloader {
	return &HTTPDownloader{
		Email:      email,
		Password:   password,
		ProjectURL: projectURL,
	}
}

// HTTPDownloader represents sharelatex project downloader
type HTTPDownloader struct {
	Email, Password, ProjectURL, Output string
}

// Download downloads project zip using credentials provided using wd as current directory
func (d *HTTPDownloader) Download(wd string) error {
	log.Printf("Downloading %s...\n", d.ProjectURL)

	projectID, isSharedProject, needAuth := parseProjectID(d.ProjectURL)
	output := fmt.Sprintf("%s.zip", projectID)

	client := resty.New()
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(1))
	client.SetOutputDirectory(wd)

	if os.Getenv("DEBUG") != "" {
		client.Debug = true
	}

	var err error
	if needAuth {
		err = setupAuth(client, d.Email, d.Password)
	}
	if isSharedProject {
		projectID, err = setupSharedProject(client, d.ProjectURL)
	}
	if err != nil {
		return err
	}

	// download the project
	projectURL := fmt.Sprintf(projectZipURLFormat, projectID)
	_, err = client.R().
		SetOutput(output).
		Get(projectURL)
	return err
}

func setupAuth(client *resty.Client, email, password string) error {
	// create inital request
	resp, err := client.R().Get(loginURL)
	if err != nil {
		return err
	}
	updateSessCookie(client, resp.Cookies())

	// get csrf
	re, err := regexp.Compile(`window.csrfToken\s=\s\"(?P<csrf>.*)\"`)
	if err != nil {
		panic(err)
	}
	matches := re.FindStringSubmatch(string(resp.Body()))
	if len(matches) != 2 {
		return fmt.Errorf("failed parsing csrf token")
	}
	csrf := matches[1]

	// do actual login
	resp, err = client.R().
		SetFormData(map[string]string{
			"_csrf":    csrf,
			"email":    email,
			"password": password,
		}).
		Post(loginURL)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		log.Println("Authentication failed.")
		return errors.New("wrong email or password")
	}

	updateSessCookie(client, resp.Cookies())
	return nil
}

func setupSharedProject(client *resty.Client, projectURL string) (string, error) {
	// create initial request
	resp, err := client.R().Get(projectURL)
	if err != nil {
		return "", err
	}

	// get real project id
	re, err := regexp.Compile(`window.project_id\s=\s\"(?P<project_id>.*)\"`)
	if err != nil {
		panic(err)
	}
	matches := re.FindStringSubmatch(string(resp.Body()))
	if len(matches) != 2 {
		return "", fmt.Errorf("failed parsing project id")
	}
	projectID := matches[1]

	updateSessCookie(client, resp.Cookies())
	return projectID, nil
}

func updateSessCookie(client *resty.Client, cookies []*http.Cookie) {
	sessCookie := getSessCookie(cookies)
	for _, cookie := range client.Cookies {
		if cookie.Name == sessionKey {
			cookie.Value = sessCookie.Value
			return
		}
	}
	client.SetCookie(sessCookie)
}

func getSessCookie(cookies []*http.Cookie) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == sessionKey {
			return cookie
		}
	}
	return nil
}
