.PHONY: go-test-coverage
go-test-coverage: ## Run go unit tests
go-test-coverage:
	go test ./... -coverprofile=coverage.xml -covermode=atomic
