TARGET=oppapi

SRC := $(wildcard *.go cmd/*/*.go internal/*/*.go pkg/*/*.go)

VERSION=$(shell git describe --tags --long --dirty 2>/dev/null)

ifeq ($(VERSION),)
	VERSION = UNKNOWN
endif

LDFLAGS=-ldflags "-X main.version=${VERSION}"

# Trigger usage of CGDEBUG env var in docker/image target
GODEBUG=gctrace=1

$(TARGET): $(SRC)
	go build $(LDFLAGS) -o ./bin/$@ ./cmd/$(TARGET)/main.go

run:
	go run ./cmd/$(TARGET)/main.go

clean:
	rm -f ./bin/$(TARGET)
	find . -name *~ -exec rm {} \;

help:
	@echo '* Help *'
	@echo
	@echo '** Common build commands **'
	@echo
	@echo 'Usage:'
	@echo '    make                       build applicatetion locally'
	@echo '    make run                   Run application locally'
	@echo '    make clean                 clean all'
	@make --no-print-directory go/help docker/help test/help swagger/help minikube/help

.PHONY: $(TARGET) run clean help

-include .env
-include build/include/include.*.mk
