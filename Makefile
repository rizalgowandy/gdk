.PHONY: help
help: # print all available make commands and their usages.
	@printf "\e[32musage: make [target]\n\n\e[0m"
	@grep -E '^[a-zA-Z_-]+:.*?# .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: # install configuration and dependencies for development.
	@./scripts/setup.sh

linter: # run linter to keep code clean
	@./scripts/linter.sh

build: # ensure all binary can be build.
	@go build -o bin/go_build_gdk
