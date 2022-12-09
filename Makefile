.PHONY: test help
.DEFAULT_GOAL:=help

test: ## run tests
	go run github.com/onsi/ginkgo/v2/ginkgo --coverprofile=coverage.txt --covermode=atomic -cover ./...

help:  ## Shows help
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'