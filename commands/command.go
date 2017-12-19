package commands

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func run(prog string, args ...string) {
	log.Printf("Running: %s %s\n", prog, strings.Join(args, " "))

	cmd := exec.Command(prog, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Run(); err == exec.ErrNotFound {
		log.Fatalf("Cannot find %s executable\n", prog)
	} else if err != nil {
		log.Fatal(strings.Join(cmd.Args, " "), ": ", err, "\n", stderr.String())
	}

	output := strings.TrimSpace(stdout.String())
	if output != "" {
		log.Println(output)
	}
}

func runWithChdir(prog, chdir string, args ...string) {
	log.Printf("Running: %s %s\n", prog, strings.Join(args, " "))

	cmd := exec.Command(prog, args...)
	cmd.Dir = chdir
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Run(); err == exec.ErrNotFound {
		log.Fatalf("Cannot find %s executable\n", prog)
	} else if err != nil {
		log.Fatal(strings.Join(cmd.Args, " "), ": ", err, "\n", stderr.String())
	}

	output := strings.TrimSpace(stdout.String())
	if output != "" {
		log.Println(output)
	}
}
