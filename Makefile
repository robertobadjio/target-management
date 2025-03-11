#!/usr/bin/make
.DEFAULT_GOAL := help
.PHONY: help

GO_TOOL_COMMAND = go tool
BOMBARDIER_COMMAND = ./bin/bombardier

install-deps-linux: ## Install dependencies for Linux
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s v1.64.5
	wget https://github.com/codesenberg/bombardier/releases/download/v2.0.2/bombardier-linux-amd64 -O ./bin/bombardier
	chmod a+x ./bin/bombardier

help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

fmt: ## Automatically format source code
	go fmt ./...
.PHONY:fmt

lint: fmt ## Check code (lint)
	./bin/golangci-lint run ./... --config .golangci.pipeline.yaml
.PHONY:lint

vet: fmt ## Check code (vet)
	go vet ./...
.PHONY:vet

vet-shadow: fmt ## Check code with detect shadow (vet)
	go vet -vettool=$(which shadow) ./...
.PHONY:vet

up: vet ## Start services
	$(DOCKER) up --build -d

load-testing: ## Run load testing
	$(BOMBARDIER_COMMAND) \
		-c 10 \
	 	-m POST http://localhost:8081/load-facts \
	 	-b "period_start=2024-12-01&period_end=2024-12-31&period_key=month&indicator_to_mo_id=227373&indicator_to_mo_fact_id=0&value=3&fact_time=2024-12-31&is_plan=0&auth_user_id=40&comment=buffer Fatkhullin" \
	 	-H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8" \
	 	-n 10 \
	 	-l