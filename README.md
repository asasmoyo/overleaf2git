# overleaf2git

[![Build Status](https://travis-ci.org/asasmoyo/overleaf2git.svg?branch=master)](https://travis-ci.org/asasmoyo/overleaf2git)

Simple script for downloading Overleaf projects and store it into Git repository

**NOTE:** This is unofficial project. Use it at your own risk!

## Usage

```
NAME:
   overleaf2git - Sync Overleaf files into git

USAGE:
   overleaf2git [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --session-key value, -s value  Your overleaf active session key
   --url value, -u value          Overleaf project url
   --workdir value, --wd value    Working directory (default: "project")
   --git-url value                Git repository url
   --git-branch value             Git repository target branch (default: "master")
   --git-force-push               Use git force push
   --help, -h                     show help
   --version, -v                  print the version
```

## Requirements

1. Unzip

2. Git

## Features

1. Supports public and private projects

2. Push project files into git repository

## Installation

You can download the binary in [release page.](https://github.com/asasmoyo/overleaf2git/releases)

Or you can download and build manually using Golang:

```
go get github.com/asasmoyo/overleaf2git/cmd/overleaf2git
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
