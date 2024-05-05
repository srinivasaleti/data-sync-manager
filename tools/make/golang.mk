.PHONY: go-test-coverage
go-test-coverage: ## Run go unit tests
go-test-coverage:
	@go test $(shell  go list ./... | grep -v 'mocks') -coverprofile=coverage.xml
	@go tool cover -html coverage.xml -o cover.html
	open cover.html

