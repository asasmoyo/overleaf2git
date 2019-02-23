.PHONY: build
build:
	mkdir -p bin
	go build -o bin/overleaf2git github.com/asasmoyo/overleaf2git/cmd/overleaf2git

.PHONY: build-ci
build-ci:
	mkdir -p bin
	env GOOS=linux GOARCH=amd64 go build -o bin/overleaf2git-linux github.com/asasmoyo/overleaf2git/cmd/overleaf2git
	env GOOS=darwin GOARCH=amd64 go build -o bin/overleaf2git-darwin github.com/asasmoyo/overleaf2git/cmd/overleaf2git

.PHONY: clean
clean:
	rm -rf bin

.PHONY: run
run:
	./bin/overleaf2git
