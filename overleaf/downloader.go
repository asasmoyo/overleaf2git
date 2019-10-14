package overleaf

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/resty.v1"
)

const (
	loginURL            = "https://www.overleaf.com/login"
	projectZipURLFormat = "https://www.overleaf.com/project/%s/download/zip"
	sessionKey          = "overleaf_session"
	output              = "project.zip"
)

// Downloader interface for downloading sharelatex project
type Downloader interface {
	Download(wd string) error
}

// NewDownloader create a new downloader instance
func NewDownloader(sessionKey, projectURL string) Downloader {
	return &HTTPDownloader{
		SessionKey: sessionKey,
		ProjectURL: projectURL,
	}
}

// HTTPDownloader represents sharelatex project downloader
type HTTPDownloader struct {
	SessionKey, ProjectURL, Output string
}

// Download downloads project zip using credentials provided using wd as current directory
func (d *HTTPDownloader) Download(wd string) error {
	log.Printf("Downloading %s\n", d.ProjectURL)

	client := resty.New()
	client.SetHeader("User-Agent", "overleaf2git - https://github.com/asasmoyo/overleaf2git")
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(3))
	client.SetOutputDirectory(wd)
	client.SetCookie(&http.Cookie{
		Name:  sessionKey,
		Value: d.SessionKey,
	})

	if os.Getenv("DEBUG") != "" {
		client.Debug = true
	}

	// download the project
	projectID := parseProjectID(d.ProjectURL)
	projectURL := fmt.Sprintf(projectZipURLFormat, projectID)
	_, err := client.R().
		SetOutput(output).
		Get(projectURL)
	return err
}
