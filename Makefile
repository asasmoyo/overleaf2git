.PHONY: build
build:
	mkdir -p bin
	go build -o bin/sharelatex2git github.com/asasmoyo/sharelatex2git/cmd/sharelatex2git

.PHONY: build-ci
	mkdir -p bin
	env GOOS=linux GOARCH=amd64 go build -o bin/sharelatex2git-linux github.com/asasmoyo/sharelatex2git/cmd/sharelatex2git
	env GOOS=darwin GOARCH=amd64 go build -o bin/sharelatex2git-darwin github.com/asasmoyo/sharelatex2git/cmd/sharelatex2git

.PHONY: deps
deps:
	go get github.com/Masterminds/glide
	go get golang.org/x/net/publicsuffix

.PHONY: clean
clean:
	rm -rf bin

.PHONY: run
run:
	./bin/sharelatex2git
