#!/usr/bin/env bash
.PHONY: help build

default: help

help:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

small: ## generate a small datafile for testing purpose
	go run generator/gen.go 1000

big: ## generate the one billion row file watch for the 13Go file coming :o
	go run generator/gen.go 1000000000

test: ## test if all impl pass the test
	go test ./...

bench: ## bench all impl
	go test -bench=. ./reader

all: ## run all impl
	go build -o app
	time ./app R1
	time ./app R2
	time ./app R3

