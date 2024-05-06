include tools/make/golang.mk
include tools/make/docker.mk

.PHONY: help
help: ## Help
	@echo "Usage: make [target]"
	@echo
	@echo "Available targets:"
	@awk '/^[a-zA-Z0-9%._-]+:.*?## (.+)$$/ { \
		printf "  %-20s %s\n", $$1, substr($$0, index($$0, "## ") + 3) \
	}' $(MAKEFILE_LIST) | sort