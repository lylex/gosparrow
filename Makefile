include .env

VERSION := $(shell git describe --tags --always --dirty)

BUILD_IMAGE ?= $(GOLANG_IMAGE_NAME):$(GOLANG_IMAGE_VERSION)

SRC_DIRS := cmd pkg
ALL_PROTOCALS := grpc rest
ROOT_PKG := $(PROJECT)
CMD_DIR := $(ROOT_PKG)/$(CMD_PKG)
BUILDABLE_DIRS := $(notdir $(patsubst %/,%, $(wildcard $(CMD_PKG)/*/)))

all: build

.PHONY: build
build: $(addprefix build-, $(BUILDABLE_DIRS))
build-%:
	@$(MAKE) TARGET=$* perform-build

.PHONY: perform-build
perform-build: bin/$(TARGET)

bin/$(TARGET):
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
	        TARGET=$(TARGET)             \
	        CMD_PKG=$(CMD_PKG)           \
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
