# disable the upto date check of make for these goals
.PHONY: build-server build-client 
.DEFAULT_GOAL := help

build-server: ## builds the rest-error-simulator server
	export GOARCH=amd64 && go build \
		-tags release \
		-o bin/rest-error-simulator-server \
		cmd/server/main.go

build-client: ## builds the rest-error-simulator client
	export GOARCH=amd64 && go build \
		-tags release \
		-o bin/rest-error-simulator-client \
		cmd/client/main.go

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
