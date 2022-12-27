GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOMOD := $(GOCMD) mod
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get

OS := $(shell uname | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m | tr '[:upper:]' '[:lower:]')
PLATFORM ?= linux/amd64

PROJECTNAME := lipsumgo
VERSION := $(shell git describe --tags --always)
BUILD := $(shell git rev-parse --short HEAD)
DOCKERIMAGE ?= $(PROJECTNAME)

GOBUILDFLAGS := -ldflags "-X=main.appName=$(PROJECTNAME) -X=main.appVersion=$(VERSION) -X=main.appBuild=$(BUILD)"
GOPATH := $(PWD)/.go

.DEFAULT_GOAL := help

.PHONY: all build clean clean-all test proto vendor vendor-clean docker help

all: clean build

build: vendor
	mkdir -p target/
	CGO_ENABLED=0 $(GOBUILD) $(GOBUILDFLAGS) -mod readonly -v -o target/lipsumgo-server ./cmd/server
	CGO_ENABLED=0 $(GOBUILD) $(GOBUILDFLAGS) -mod readonly -v -o target/lipsumgo-client ./cmd/client
	# build ok

clean:
	rm -rf target/
	# clean ok

clean-all:
	rm -rf target/
	$(GOCLEAN) -cache -testcache -modcache ./...
	# clean-all ok

test:
	$(GOTEST) $(GOBUILDFLAGS) -mod readonly -v ./...
	# test ok

proto:
	proto/gen.sh
	# proto ok

vendor:
	$(GOMOD) download
	$(GOMOD) vendor
	$(GOMOD) verify
	# vendor ok

vendor-clean:
	rm -rf vendor/
	chmod -R 700 "$(PWD)/.go" && rm -rf "$(PWD)/.go"
	# vendor-clean ok

docker:
	docker buildx build \
	  --platform "$(PLATFORM)" \
	  --progress plain \
	  -f Dockerfile \
	  -t "$(DOCKERIMAGE):$(BUILD)" \
	  -t "$(DOCKERIMAGE):$(VERSION)" \
	  .
	# docker ok

help: Makefile
	@echo "To make \"$(PROJECTNAME)\", use one of the following commands:"
	@echo "    all"
	@echo "    build"
	@echo "    clean"
	@echo "    clean-all"
	@echo "    test"
	@echo "    proto"
	@echo "    vendor"
	@echo "    vendor-clean"
	@echo "    docker"
	@echo "    help"
	@echo
