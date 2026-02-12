.PHONY: all tidy fa fmt lint test

all: tidy fa fmt lint

tidy:
	go mod tidy

fa:
	@fieldalignment -fix ./...

fmt:
	@goimports -w -local github.com/pixel365/goepp .
	@gofmt -w .
	@golines -w .

lint:
	@golangci-lint run

test:
	@go test ./...


