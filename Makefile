PROJECT := gosparrow
VERSION := $(shell git describe --tags --always --dirty)

DOCKER_REGISTRY ?= 127.0.0.1:5000/local
IMAGE := $(DOCKER_REGISTRY)/$(PROJECT)
BUILD_IMAGE ?= golang:1.8

SRC_DIRS := cmd pkg
ROOT_PKG := $(PROJECT)

all: build

build: bin/$(PROJECT)

bin/$(PROJECT):
	@echo "Building $@ in container..."
	@docker run                          \
	    --rm                             \
	    -v "$$(pwd):/go/src/$(ROOT_PKG)" \
	    -v "$$(pwd)/.go:/go"             \
	    -v "$$(pwd)/bin:/go/bin"         \
	    -w /go/src/$(ROOT_PKG)           \
	    $(BUILD_IMAGE)                   \
	    /bin/sh -c "                     \
	        VERSION=$(VERSION)           \
	        ./build/build.sh             \
	    "

.PHONY: clean
clean:
	@docker run                          \
	    --rm                             \
	    -v "$$(pwd):/go/src/$(ROOT_PKG)" \
	    -w /go/src/$(ROOT_PKG)           \
	    $(BUILD_IMAGE)                   \
	    /bin/sh -c "                     \
	        rm -f bin/*                  \
	    "
