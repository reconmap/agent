SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.ONESHELL:
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
LATEST_TAG = $(shell git describe --tags)

CONTAINER_NAME=quay.io/reconmap/agent:latest

PROGRAM=reconmapd

$(PROGRAM):
	go build -v -ldflags="-X 'github.com/reconmap/agent/internal/build.BuildVersion=$(LATEST_TAG)'" -o $(PROGRAM) ./cmd/reconmapd

.PHONY: tests
tests:
	go test ./...

.PHONY: clean
clean:
	rm -f $(PROGRAM)

.PHONY: docker-build
docker-build:
	docker build -t $(CONTAINER_NAME) .

.PHONY: docker-push
docker-push:
	docker push $(CONTAINER_NAME)

.PHONY: lint
lint: GOLANGCI_LINT_VERSION ?= 1.49
lint:
	docker run \
	-v $(CURDIR):/reconmap/agent \
	-w /reconmap/agent \
	golangci/golangci-lint:v$(GOLANGCI_LINT_VERSION)-alpine \
	golangci-lint run -c .golangci.yml --timeout 10m --fix

.PHONY: update-deps
update-deps:
	go get -u ./...

