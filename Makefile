SHELL := bash

ifndef NO_COLOR
YELLOW=\033[0;33m
CYAN=\033[1;36m
RED=\033[31m
# no color
NC=\033[0m
endif

GOLIC_VERSION  ?= v0.7.2
GOLINT_VERSION  ?= v1.45.0


test:
	go test ./... --cover

.PHONY: lint
lint:
	goimports -w ./
	@echo -e "\n$(YELLOW)Running the linters$(NC)"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLINT_VERSION)
	$(GOBIN)/golangci-lint run


check: lint test

all: check
