ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

VERSION := $(shell git tag | grep ^v | sort -V | tail -n 1)
LDFLAGS = -ldflags "-X main.version=${VERSION}"

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## creates binary
	go build ${LDFLAGS} .
install: build ## compiles and installs into system
	rm ~/.magento-cloud/bin/magento
	cp ./magento-cli ~/.magento-cloud/bin/magento
	sudo chmod +x ~/.magento-cloud/bin/magento
run: ## run the command through go, accepts args i.e. `make run -- build -h`
	go run ${LDFLAGS} ./main.go $(RUN_ARGS)
test: build
	go test --cover ./... 
coverage: build ## run test suite suitable for codecov.io
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...
lcov: coverage ## lcov formatted test files for ide interpretation
	mkdir -p coverage
	gcov2lcov -infile=coverage.txt -outfile=coverage/lcov-vscode.info
lcov-deps: ## install dependancies for gcov2lcov
	go get -u github.com/jandelgado/gcov2lcov
docs-install: ## installs docs dependancies
	cd docs/ && npm install
docs-build: docs-install ## builds the docs instance
	cd docs/ && hugo --minify
docs-serve: docs-install ## builds and serves docs
	cd docs/ && npx hugo serve
