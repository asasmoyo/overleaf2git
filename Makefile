.PHONY: build
build:
	mkdir -p bin
	go build -o bin/sharelatex2git github.com/asasmoyo/sharelatex2git/cmd/sharelatex2git

.PHONY: clean
clean:
	rm -rf bin

.PHONY: run
run:
	./bin/sharelatex2git
