.PHONY: test build init

all: build

build: test
	@echo "Building"
	@go build -o bin/ssb src/github.com/Nueard/ssb/main.go

test: init
	@echo "Running unit tests"
	@go test github.com/Nueard/ssb/...

init:
	@echo "Ensuring dependencies"
	@cd src/github.com/Nueard/ssb; dep ensure

run: test
	@go run src/github.com/Nueard/ssb/main.go