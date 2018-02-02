# Makefile configuration
.DEFAULT_GOAL := help
.PHONY: fmt vet test cyclo lint

ok: fmt vet test cyclo lint ## Checks project

fmt: ## Golang code formatting tool
	@gofmt -s -w .

vet: ## Check code against common errors
	@go vet ./...

test: ## Run tests
	@go test ./...

cyclo: ## Check cyclomatic complexity
	@${GOPATH}/bin/gocyclo -over 15 .

lint: ## Code linting
	@${GOPATH}/bin/golint ./...

deps: ## Download required dependencies
	@go get ./...

build: ## Builds telegacli
	@GOOS="linux" GOARCH="amd64" go build -o bin/telegacli main.go

help:
	@grep --extended-regexp '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
