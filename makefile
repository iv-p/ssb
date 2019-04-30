.PHONY: test build

build: test
	go build -o bin/ssb main.go

test:
	go test ./...