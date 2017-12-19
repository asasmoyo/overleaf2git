package sharelatex

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-resty/resty"
)

const (
	sharelatexLoginURL         = "https://www.sharelatex.com/login"
	sharelatexProjectURLFormat = "https://www.sharelatex.com/project/%s/download/zip"
	sharelatexSessKey          = "sharelatex_session"
)

// Downloader interface for downloading sharelatex project
type Downloader interface {
	Download(wd string) error
}

// NewDownloader create a new downloader instance
func NewDownloader(email, password, projectID, output string, isPublic bool) Downloader {
	return &HTTPDownloader{
		Email:     email,
		Password:  password,
		ProjectID: projectID,
		IsPublic:  isPublic,
		Output:    output,
	}
}

// HTTPDownloader representds sharelatex project downloader
// It expects the combination of provided Email, Password and ProjectID are valid
// Authentication will be skipped if IsPublic is true
type HTTPDownloader struct {
	Email, Password, ProjectID, Output string
	IsPublic                           bool
}

// Download downloads project zip using credentials provided using wd as current directory
func (d *HTTPDownloader) Download(wd string) error {
	output := d.Output
	if output == "" {
		output = fmt.Sprintf("%s.zip", d.ProjectID)
	}

	client := resty.New()
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(1))
	client.SetOutputDirectory(wd)

	// get initial cookies
	resp, err := client.R().Get(sharelatexLoginURL)
	if err != nil {
		return err
	}

	// get csrf
	re, err := regexp.Compile(`window.csrfToken\s=\s\"(?P<csrf>.*)\"`)
	if err != nil {
		panic(err)
	}
	matches := re.FindStringSubmatch(string(resp.Body()))
	if len(matches) != 2 {
		return fmt.Errorf("cannot get csrf")
	}
	csrf := matches[1]

	// do actual login
	updateSessCookie(client, getSessCookie(resp.Cookies()))
	resp, err = client.R().SetFormData(map[string]string{
		"_csrf":    csrf,
		"email":    d.Email,
		"password": d.Password,
	}).Post(sharelatexLoginURL)
	if err != nil {
		return err
	}

	// download the project
	updateSessCookie(client, getSessCookie(resp.Cookies()))
	projectURL := fmt.Sprintf(sharelatexProjectURLFormat, d.ProjectID)
	resp, err = client.R().
		SetOutput(output).
		Get(projectURL)
	return err
}

func getSessCookie(cookies []*http.Cookie) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == sharelatexSessKey {
			return cookie
		}
	}
	return nil
}

func updateSessCookie(client *resty.Client, sessCookie *http.Cookie) {
	for _, cookie := range client.Cookies {
		if cookie.Name == sharelatexSessKey {
			cookie.Value = sessCookie.Value
			return
		}
	}
	client.SetCookie(sessCookie)
}
