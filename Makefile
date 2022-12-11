.PHONY: test gernerate build docker-build help
.DEFAULT_GOAL:=help

GOARCH?=$(shell go env GOARCH)
GOARM?=$(shell go env GOARM)

test: ## run tests
	go run github.com/onsi/ginkgo/v2/ginkgo --coverprofile=coverage.txt --covermode=atomic -cover ./...

gernerate: ## run go generate
	go generate ./..

build:  ## Build binary
	go build -v -ldflags="-w -s" -o /tmp/gss

docker-build:  ## Build docker image 
	docker buildx build \
		--network=host \
		-o type=docker \
		-t ghcr.io/kwitsch/gosimpleserve \
		.

help:  ## Shows help
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'