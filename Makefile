TAG=$(shell git describe --tags --abbrev=0 | sed -Ee 's/^v|-.*//')
COMMIT=$(shell git rev-list -1 HEAD)
DATE=$(shell date '+%Y-%m-%d %H:%M:%S')
GITTAGHEAD=$(shell git describe --tags --abbrev=0 @^)

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

start: ## start docker-compose
	docker-compose up

build: ## build docker-compose
	docker-compose --build