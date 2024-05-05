.PHONY: go-test-coverage
go-test-coverage: ## Run go unit tests
go-test-coverage:
	@go test $(shell  go list ./... | grep -v 'mocks') -coverprofile=coverage.xml -covermode=atomic

