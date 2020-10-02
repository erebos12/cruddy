.DEFAULT_GOAL:=local
SHELL:=/bin/bash

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\033[36m***This is the Makefile to start the ESRA MADE locally***\033[0m \n \nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

kill: ## kill'em all
	docker-compose down

local: kill ## start all locally
	docker-compose up --build

purge: ## Remove dangling images (have a look with cat docker_purge.sh)
	./docker_purge.sh 


