help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo "Please use \`make <ROOT>' where <ROOT> is one of"
	@echo "  clean				to remove previous version binaries"
	@echo "  all				to test and build"
	@echo "  build				to clean, make dependencies and make tange"
	@echo "  dependencies			to install the dependencies"
	@echo "  tange			to install vendors and build the main binary"
	@echo "  serve				to make tange and run with serve arg"
	@echo "  build-linux			to build the main binary for Ubuntu"
	@echo "  test				to run unittests"
	@echo "  docker-all			to docker-build + docker-push"
	@echo "  docker-build			to build docker images in registry and tag it"
	@echo "  docker-run			to run the tange docker image"
	@echo "  docker-push			to push current version of docker image to the registry"
	@echo "  install-dep		to install dep (package versions)"
	@echo "  action <args>		to build & run binary with <args>"

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
all: test build
build: clean tange
tange:
	go install ./vendor/...
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
serve:
	$(GOBUILD) -o $(BINARY_NAME) #-v ./...
	./$(BINARY_NAME) serve
update-dependencies:
	dep ensure -update
dependencies: install-dep
	dep ensure

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

docker-all: docker-build docker-push 

docker-build:
	docker build -t $(DOCKER_IMAGE):$(VERSION) .
	docker tag $(DOCKER_IMAGE):$(VERSION) $(DOCKER_IMAGE):latest

docker-run:
	docker run --rm -it -p $(PORT):$(PORT) $(DOCKER_IMAGE)

docker-push:
	docker push $(DOCKER_IMAGE)

install-dep:
	dep version 2> /dev/null || go get -u github.com/golang/dep/cmd/dep
action: build
	echo args=$(args)
	./$(BINARY_NAME) $(args)

PORT=61613

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=tange
BINARY_UNIX=$(BINARY_NAME)_unix
ROOT := tange 
DOCKER_IMAGE := tange 
PROJECT_NAME := "tange"
PKG := "$(PROJECT_NAME)"
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)


GO_VARS ?=
GO ?= CGO_ENABLED=1 go
GIT ?= git
COMMIT := $(shell $(GIT) rev-parse HEAD)
VERSION ?= $(shell $(GIT) describe --tags ${COMMIT} 2> /dev/null || echo "$(COMMIT)")
BUILD_TIME := $(shell LANG=en_US date +"%F_%T_%z")
LD_FLAGS := -X $(ROOT).Version=$(VERSION) -X $(ROOT).Commit=$(COMMIT) -X $(ROOT).BuildTime=$(BUILD_TIME)
UNIX_TIME := $(shell date +%s)

