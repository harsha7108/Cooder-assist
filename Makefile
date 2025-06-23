.PHONY: default buildall buildclient buildclient compose test testit check fmt lint fix compose

GO 			 := /usr/local/bin/go
BINDIR       := $(CURDIR)/bin
BINNAME_BASE := cooder-assist
MAINDIR_BASE := $(CURDIR)/cmd
SRC          := $(shell find . -type f -name '*.go' -print)
LDFLAGS      := -s

SHELL      = /usr/bin/env bash

GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

GOBIN		 := $(shell which go)

ifneq ("$(wildcard $(GOBIN))","")
	GO = $(GOBIN)
endif

ifdef VERSION
	BINARY_VERSION = $(VERSION)
endif
BINARY_VERSION ?= ${GIT_SHA}
BUILD_DATE = $(shell date +'%F')
BUILD_TIME = $(shell date +'%T')  #removed %z as its causing issues with flags

LDFLAGS += -X main.version=${BINARY_VERSION}
LDFLAGS += -X main.gitCommit=${GIT_COMMIT}
LDFLAGS += -X main.gitTreeState=${GIT_DIRTY}
LDFLAGS += -X main.buildDate=${BUILD_DATE}
LDFLAGS += -X main.buildTime=${BUILD_TIME}


default: fmt lint buildlinux build

build: $(SRC)
	@echo build cooder-assist binary
	
	$(GO) build -v -ldflags '$(LDFLAGS)' -o $(BINDIR)/$(BINNAME_BASE)-local $(MAINDIR_BASE)


buildlinux: $(SRC)
	GOOS=linux GOARCH=amd64 $(GO) build -v -ldflags '$(LDFLAGS)' -o $(BINDIR)/$(BINNAME_BASE) $(MAINDIR_BASE)

buildcontainer: buildlinux
	docker build -t $(BINNAME_BASE):${BINARY_VERSION} .

fmt:
	@echo go fmt ./...
	$(GO) fmt ./...

lint:
	@echo run golangci-lint on project
	@golangci-lint run --allow-parallel-runners ./...
	
clean:
	@rm -rf $(BINDIR)