# sharelatex2git

[![Build Status](https://travis-ci.org/asasmoyo/sharelatex2git.svg?branch=master)](https://travis-ci.org/asasmoyo/sharelatex2git)

Simple script for downloading Sharelatex projects and store it into Git repository

**NOTE:** This is unofficial project. Use it at your own risk!

## Usage

```
NAME:
   sharelatex2git - Sync sharelatex files into git

USAGE:
   sharelatex2git [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --email value, -e value      Your sharelatex email
   --password value, -p value   Your sharelatex password
   --url value, -u value        Sharelatex project url
   --workdir value, --wd value  Working directory (default: "project")
   --git-url value              Git repository url
   --git-branch value           Git repository target branch (default: "master")
   --git-force-push             Use git force push
   --help, -h                   show help
   --version, -v                print the version
```

## Requirements

1. Unzip

2. Git

## Features

1. Supports public and private projects

2. Push project files into git repository

## Installation

You can download the binary in [release page.](https://github.com/asasmoyo/sharelatex2git/releases)

Or you can download and build manually using Golang:

```
go get github.com/asasmoyo/sharelatex2git/cmd/sharelatex2git
```

## Development

Requirements:

1. Make

2. Golang (let's use latest version)

You can install required dependencies by:

```
make deps
```

## Licence

MIT Licence