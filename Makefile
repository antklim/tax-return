.PHONY: build
build: go-build ## Build Go binaries

.PHONY: clean
clean: go-clean ## Clean build cache and dependencies

.PHONY: deps
deps: go-deps ## Install dependencies

.PHONY: test
test: go-test ## Run tests

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

go-build:
	@rm -rf build
	@mkdir build
	go build -o build -v ./...

go-clean: go-clean-cache go-clean-deps

go-clean-cache:
	go clean -cache

go-clean-test-cache:
	go clean -testcache

go-clean-deps:
	go mod tidy

go-deps:
	go mod download

go-test:
	go test -v ./...

.DEFAULT_GOAL := help