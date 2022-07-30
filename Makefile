# Environment
GOPATH ?= $(HOME)/go
BIN_DIR = $(GOPATH)/bin
TMPDIR ?= $(shell dirname $$(mktemp -u))

# Project specific variables
COVER_FILE ?= coverage.out

# Tools
.PHONY: tools
tools: ## Install all needed tools, e.g. for static checks
	@echo Installing tools from req-tools.txt
	@grep '@' req-tools.txt | xargs -tI % go install %


# Main targets
.PHONY: test
test: ## Run unit (short) tests
	go test -short ./... -coverprofile=$(COVER_FILE)
	go tool cover -func=$(COVER_FILE) | grep ^total

$(COVER_FILE):
	$(MAKE) test

.PHONY: cover
cover: $(COVER_FILE) ## Output coverage in human readable form in html
	go tool cover -html=$(COVER_FILE)
	rm -f $(COVER_FILE)


.PHONY: lint
lint: tools ## Check the project with lint
	staticcheck ./...

.PHONY: vet
vet: ## Check the project with vet
	go vet ./...

.PHONY: fmt
fmt: ## Run go fmt for the whole project
	test -z $$(for d in $$(go list -f {{.Dir}} ./...); do gofmt -e -l -w $$d/*.go; done)

.PHONY: imports
imports: $(GOIMPORTS) ## Check and fix import section by import rules
	test -z $$(for d in $$(go list -f {{.Dir}} ./...); do goimports -e -l -local $$(go list) -w $$d/*.go; done)

.PHONY: cyclomatic
cyclomatic: tools ## Check the project with gocyclo for cyclomatic complexity
	gocyclo -over 10 `find . -type f -iname '*.go' ! -iname "*_test.go" -not -path '*/\.*'`

.PHONY: static_check
static_check: fmt imports vet lint cyclomatic ## Run static checks (fmt, lint, imports, vet, ...) all over the project

.PHONY: check
check: static_check test ## Check project with static checks and unit tests

.PHONY: dependencies
dependencies: ## Manage go mod dependencies, beautify go.mod and go.sum files
	go mod tidy

.PHONY: clean
clean: ## Clean the project from built files
	rm -f $(COVER_FILE)

.PHONY: help
help: ## Print this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
