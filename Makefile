# Makefile for go-validator

SHELL := /bin/bash
GO ?= go
PKG ?= ./...
TEST_FLAGS ?= -race -count=1
COVER_PROFILE ?= coverage.out
BIN ?= go-validator
GOLANGCI_LINT ?= golangci-lint
LINT_VERSION ?= v1.61.0

.PHONY: help lint tidy test coverage bench clean

help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"}; /^[a-zA-Z0-9_.-]+:.*##/ {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Run unit tests
	$(GO) test $(TEST_FLAGS) $(PKG)

coverage: ## Run tests with coverage and generate reports
	$(GO) test -race -covermode=atomic -coverprofile $(COVER_PROFILE) $(PKG)
	$(GO) tool cover -func=$(COVER_PROFILE) | tail -n 1
	@echo "HTML report: $(COVER_PROFILE).html"
	$(GO) tool cover -html=$(COVER_PROFILE) -o $(COVER_PROFILE).html

bench: ## Run benchmarks
	$(GO) test -bench=. -benchmem $(PKG)

lint: ## Run golangci-lint
	$(GOLANGCI_LINT) run ./...

clean: ## Clean generated artifacts
	rm -f $(COVER_PROFILE) $(COVER_PROFILE).html
