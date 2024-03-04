#!/usr/bin/env bash
.PHONY: help build

default: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

small:
	go run generator/gen.go 1000

big:
	go run generator/gen.go 1000000000

test:
	go test ./...

bench:
	go test -bench=. ./reader

r1:
	go run main.go R1

