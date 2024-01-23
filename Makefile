-include .env

GITHASH = $(GITHUB_SHA)
ifndef GITHUB_SHA
		GITHASH = $(shell git rev-list -1 HEAD)
endif

VERSION = $(GITHUB_REF)
ifndef GITHUB_REF
		VERSION = $(shell git describe --tags)
endif

# Use linker flags to provide version/build settings
# LDFLAGS=-ldflags "\
#         -X 'github.com/cmd/app/main.GitHash=$(GITHASH)' \
#         -X 'github.com/cmd/app/main.Version=$(VERSION)'"


# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.DEFAULT_GOAL := default
configure:
	go mod tidy

mocks:
	go generate ./...

test:
	go test -tags=unit ./... -race

test-e2e:
	go test -tags=e2e ./... -race

coverage:
	gocov test ./... | gocov-xml > coverage.xml

build:
	cd cmd/app/ ; \
	go build $(LDFLAGS) -o ../../bin/app .

