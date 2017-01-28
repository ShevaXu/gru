GO ?= go
BUILD := $(CURDIR)/gru

all: dev

dev:
	$(GO) build

prod:
	$(GO) build --ldflags '-extldflags "-lm -lstdc++ -static"' -o $(BUILD)

deps:
	godep save ./...