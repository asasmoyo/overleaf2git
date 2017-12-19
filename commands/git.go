package commands

import (
	"fmt"
	"time"
)

func GitCloneRepo(url, dest string) {
	run("git", "clone", url, dest)
}

func GitAddAll(repoDir string) {
	runWithChdir("git", repoDir, "add", "--all")
}

func GitCommit(repoDir string) {
	msg := fmt.Sprintf("Committed with sharelatex2git at %s", time.Now().Format(time.RFC3339))
	runWithChdir("git", repoDir, "commit", "-m", fmt.Sprintf("\"%s\"", msg))
}

func GitPush(repoDir, branch string, force bool) {
	if force {
		runWithChdir("git", repoDir, "push", "-f", "origin", "master")
	} else {
		runWithChdir("git", repoDir, "push", "origin", "master")
	}
}
