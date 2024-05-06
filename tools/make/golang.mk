# This is a wrapper to build golang binaries
#
# All make targets related to golang are defined in this file.

# Supported Platforms for building multiarch binaries.
PLATFORMS ?= darwin_amd64 darwin_arm64 linux_amd64 linux_arm64 
CONFIG ?= 

.PHONY: test
test: ## Run go unit tests
test:
	@go test ./...

.PHONY: test-coverage
test-coverage:  ## Run go test coverage
test-coverage:
	@go test $(shell go list ./... | grep -v 'mocks') -coverprofile=coverage.xml

.PHONY: build
build:
build: ## Run Go Build
	@echo "building application, please wait..."
	@go build -o bin/data-sync-manager ./orchestrator

.PHONY: run
run: build
run: ## Run application
	@echo "running application..."
	@bin/data-sync-manager --config $(CONFIG)
