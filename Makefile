BINARY  ?= roaw
GOCMD   ?= go
PARALL  ?= $(shell { nproc --all || echo 1 ; } | xargs -I{} expr {} / 2 + 1 )
GOTEST  ?= $(GOCMD) test -timeout 10s -parallel $(PARALL)
GOVET   ?= $(GOCMD) vet
GOFMT   ?= gofmt
VERSION ?= $(shell git describe --tags --dirty --match='v*' 2> /dev/null || echo "$${DOCKER_BUILD_BIN_VERSION:-dev}")
COMMIT  ?= $(shell git rev-parse --short HEAD 2> /dev/null || echo "")
DATEUTC ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
FILES   ?= $(shell find . -type f -name '*.go')

GREEN  := $(shell tput -Txterm setaf 2 2>/dev/null)
YELLOW := $(shell tput -Txterm setaf 3 2>/dev/null)
WHITE  := $(shell tput -Txterm setaf 7 2>/dev/null)
CYAN   := $(shell tput -Txterm setaf 6 2>/dev/null)
RESET  := $(shell tput -Txterm sgr0    2>/dev/null)

.PHONY: all test build coverage contributors scripts

default: help

## Build:
build: ## Build your project and put the output binary in build/
	$(GOCMD) build -ldflags "-s -w -X 'main.version=$(VERSION)' -X 'main.commit=$(COMMIT)' -X 'main.dateStr=$(DATEUTC)'" -o build/$(BINARY) .

clean: ## Remove build related file
	rm -fr ./build
	rm -fr ./coverage

fmt: ## Format your code with gofmt
	$(GOFMT) -w .

## Test:
test: ## Run the tests of the project (fastest)
	$(GOVET) ./...
	$(GOTEST) ./...

test-ci: ## Run ALL the tests of the project (+race)
	$(GOVET) ./...
	$(GOTEST) -v -race ./...

test-coverage: ## Run the tests of the project and export the coverage
	rm -fr coverage && mkdir coverage
	$(GOTEST) -cover -covermode=atomic -coverprofile=coverage/coverage.out ./...
	@echo ""
	$(GOCMD) tool cover -func=coverage/coverage.out
	@echo ""
	$(GOCMD) tool cover -func=coverage/coverage.out -o coverage/coverage.txt
	$(GOCMD) tool cover -html=coverage/coverage.out -o coverage/coverage.html

coverage: test-coverage  ## Run test-coverage and open coverage in your browser
	$(GOCMD) tool cover -html=coverage/coverage.out

## Lint:
lint: lint-go ## Run all available linters

lint-go: ## Use gofmt and staticcheck on your project
ifneq (, $(shell $(GOFMT) -l . ))
	@echo "This files are not gofmt compliant:"
	@$(GOFMT) -l .
	@echo "Please run '$(CYAN)make fmt$(RESET)' to format your code"
	@exit 1
endif
ifeq (, $(shell which staticcheck 2>/dev/null))
	go install honnef.co/go/tools/cmd/staticcheck@latest
endif
	staticcheck ./...


## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)

env:    ## Print useful environment variables to stdout
	@echo '$$(GOCMD)   :' $(GOCMD)
	@echo '$$(GOTEST)  :' $(GOTEST)
	@echo '$$(GOVET)   :' $(GOVET)
	@echo '$$(BINARY)  :' $(BINARY)
	@echo '$$(VERSION) :' $(VERSION)
	@echo '$$(COMMIT)  :' $(COMMIT)
	@echo '$$(DATEUTC) :' $(DATEUTC)
	@echo '$$(FILES#)  :' $(shell echo $(FILES) | wc -w)

setup: ## Setup some dev dependencies (eg: pre-commit)
ifeq (, $(shell which pre-commit 2>/dev/null))
	@echo "pre-commit is not installed. Check $(CYAN)https://pre-commit.com/#install$(RESET)"
endif
	pre-commit install --install-hooks
