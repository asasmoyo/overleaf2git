package main

import (
	"log"
	"os"

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
			Name:  "project-id, pid",
			Usage: "Sharelatex project id",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "Output file",
		},
	}
	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) {
	projectID := c.String("project-id")
	if projectID == "" {
		log.Panicln("ProjectID cannot be empty")
	}

	downloader := sharelatex.NewDownloader(
		c.String("email"),
		c.String("password"),
		projectID,
		c.String("output"),
		false,
	)
	err := downloader.Download("")
	if err != nil {
		panic(err)
	}
}
