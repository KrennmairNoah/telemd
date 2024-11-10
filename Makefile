# Go parameters
GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

CURDIR=$(shell pwd)
export GOBIN := $(CURDIR)/bin

all:  build-all

build-all:
	$(GOINSTALL) ./...

telemd:
	$(GOINSTALL) ./cmd/telemd

clean:
	$(GOCLEAN)
	rm -rf bin/

docker:
	scripts/docker-build.sh
	#scripts/docker-release.sh
