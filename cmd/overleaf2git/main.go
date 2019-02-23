package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/asasmoyo/overleaf2git/commands"
	"github.com/asasmoyo/overleaf2git/overleaf"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "overleaf2git"
	app.Usage = "Sync overleaf files into git"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "session-key, s",
			Usage: "Your overleaf active session key",
		},
		cli.StringFlag{
			Name:  "url, u",
			Usage: "Overleaf project url",
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
	if !strings.HasPrefix(projectURL, "https://www.overleaf.com/") {
		log.Println("Project url must begin with https://www.overleaf.com/")
		return
	}

	// prepare workdir
	workdir := c.String("workdir")
	prepareWorkdir(workdir, c.String("git-url"))
	log.Printf("Using %s as workdir\n", workdir)

	downloader := overleaf.NewDownloader(
		c.String("session-key"),
		projectURL,
	)
	err := downloader.Download(workdir)
	if err != nil {
		log.Fatalln("Got an error:", err.Error())
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
