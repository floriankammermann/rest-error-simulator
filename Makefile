# disable the upto date check of make for these goals
.PHONY: build 

build: ## builds the rest-error-simulator
	export GOARCH=amd64 && go build \
		-tags release \
		-o bin/rest-error-simulator 

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
