SHELL := /bin/bash
INSTALL_DIR ?= /usr/bin
PLUGIN_BIN ?= {{ .PluginName }}
PLUGIN_DEPENDENCIES := $(shell find . -name "*.go")
# if you want to execute gotest with verbosity, set this flag to `true`.
TEST_VERBOSE ?= false

build: format test $(PLUGIN_BIN)

format:
	go fmt ./...

test:
ifeq ($(TEST_VERBOSE), true)
	go test -v ./...
else
	go test ./...
endif

$(PLUGIN_BIN): $(PLUGIN_DEPENDENCIES) generate
	go build -o $(PLUGIN_BIN) ./cmd/$(PLUGIN_BIN)/main.go

generate:
	kubectl-plugin-builder generate

install: $(PLUGIN_BIN)
	mv $(PLUGIN_BIN) $(INSTALL_DIR)