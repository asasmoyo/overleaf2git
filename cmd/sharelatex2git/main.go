package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/asasmoyo/sharelatex2git/commands"
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
			Value: "project",
		},
		cli.StringFlag{
			Name:  "git-url",
			Usage: "Git repository url",
		},
		cli.StringFlag{
			Name:  "git-branch",
			Usage: "Git repository target branch",
			Value: "master",
		},
		cli.BoolFlag{
			Name:  "git-force-push",
			Usage: "Use git force push",
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

	// prepare workdir
	workdir := c.String("workdir")
	prepareWorkdir(workdir, c.String("git-url"))
	log.Printf("Using %s as workdir\n", workdir)

	downloader := sharelatex.NewDownloader(
		c.String("email"),
		c.String("password"),
		projectURL,
	)
	err := downloader.Download(workdir)
	if err != nil {
		log.Println("Got an error:", err.Error())
	}

	projectZip := fmt.Sprintf("%s/project.zip", workdir)
	projectDir := fmt.Sprintf("%s/project", workdir)
	repoDir := fmt.Sprintf("%s/repo", workdir)

	commands.Unzip(projectZip, projectDir)
	commands.AddFiles(projectDir, repoDir)
	commands.GitAddAll(repoDir)
	commands.GitCommit(repoDir)
	commands.GitPush(repoDir, c.String("git-branch"), c.Bool("git-force-push"))
}

func prepareWorkdir(wd, gitURL string) {
	os.RemoveAll(wd)
	os.MkdirAll(wd, os.ModePerm)

	repoDir := fmt.Sprintf("%s/repo", wd)
	log.Printf("Cloning %s repo into %s\n", gitURL, repoDir)
	commands.GitCloneRepo(gitURL, repoDir)
}
