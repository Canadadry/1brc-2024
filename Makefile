#!/usr/bin/env bash
.PHONY: help build

default: help

help:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## test if all impl pass the test
	go test ./...

bench: ## bench all impl
	go test -bench=. ./reader

all: ## run all impl
	rm -f profile00*
	$(MAKE) initDocker
	$(MAKE) measureDocker SAMPLE=1000000
	$(MAKE) profilDocker READER=R1
	$(MAKE) profilDocker READER=R2
	$(MAKE) profilDocker READER=R3

initDocker:
	docker build -t 1brc-2024 .

measureDocker:
	docker run -v $$(pwd):/app 1brc-2024 bash -c "go run generator/gen.go $(SAMPLE)"

profilDocker:
	docker run -v $$(pwd):/app 1brc-2024 bash -c "go run read.go $(READER) > /dev/null"
	docker run -v $$(pwd):/app 1brc-2024 bash -c "echo \"png\" | go tool pprof cpu.profil"
