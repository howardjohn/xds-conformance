.DEFAULT_GOAL := build

SHELL := /bin/bash
BINDIR := bin
PKG := github.com/envoyproxy/xds-conformance

.PHONY: build
build:
	@go build ./...

.PHONY: test
test:
	@go test ./...

.PHONY: format
format:
	@goimports -local $(PKG) -w -l pkg
