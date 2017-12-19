package main

import (
	"log"
	"os"
	"strings"

	"github.com/asasmoyo/sharelatex2git/sharelatex"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sharelatex2git"
	app.Usage = "Sync sharelatex files into git"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "email, e",
			Usage: "Your sharelatex email",
		},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "Your sharelatex password",
		},
		cli.StringFlag{
			Name:  "url, u",
			Usage: "Sharelatex project url",
		},
		cli.StringFlag{
			Name:  "workdir, wd",
			Usage: "Working directory",
		},
	}
	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) {
	// validate project url
	projectURL := c.String("url")
	if projectURL == "" {
		log.Println("Project url cannot be empty")
		return
	}
	if !strings.HasPrefix(projectURL, "https://www.sharelatex.com/") {
		log.Println("Project url must begin with https://www.sharelatex.com/")
		return
	}

	// make sure workdir is present
	workdir := c.String("workdir")
	os.MkdirAll(workdir, os.ModePerm)

	downloader := sharelatex.NewDownloader(
		c.String("email"),
		c.String("password"),
		projectURL,
	)
	err := downloader.Download(workdir)
	if err != nil {
		log.Println("Got an error:", err.Error())
	}
}
