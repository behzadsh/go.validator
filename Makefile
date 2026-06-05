# Colors for output
GREEN  := \033[32m
YELLOW := \033[33m
CYAN   := \033[36m
RESET  := \033[0m

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help message
	@echo "$(GREEN)Go Validator Makefile$(RESET)"
	@echo "Usage: make $(CYAN)<target>$(RESET)"
	@awk 'BEGIN {FS = ":.*##";} \
		/^[a-zA-Z0-9_-]+:.*?##/ { printf "  $(CYAN)%-15s$(RESET) %s\n", $$1, $$2 } \
		/^##@/ { printf "\n$(YELLOW)%s$(RESET)\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: test test-coverage clean fmt lint tidy

test: ## Run tests
	@echo "$(GREEN)Running tests (with race)...$(RESET)"
	@go test -v -race ./...

test-coverage: ## Run tests with coverage
	@echo "$(GREEN)Running tests with coverage...$(RESET)"
	@go test -coverprofile=coverage.out ./...
	@echo "$(GREEN)Coverage file created: coverage.out$(RESET)"

clean: ## Clean coverage files
	@echo "$(YELLOW)Cleaning coverage files...$(RESET)"
	@rm -f coverage.out
	@echo "$(GREEN)Clean complete$(RESET)"

fmt: ## Format code using golangci-lint formatter
	@echo "$(GREEN)Formatting code with golangci-lint...$(RESET)"
	@golangci-lint fmt
	@echo "$(GREEN)Code formatting complete$(RESET)"

lint: ## Lint code using golangci-lint linter
	@echo "$(GREEN)Linting code with golangci-lint...$(RESET)"
	@golangci-lint run
	@echo "$(GREEN)Linting complete$(RESET)"

tidy: ## Run go mod tidy
	@echo "$(GREEN)Running go mod tidy...$(RESET)"
	@go mod tidy
	@echo "$(GREEN)Go modules tidied$(RESET)"
