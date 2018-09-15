APP = caldera
RELEASE ?= v0.0.1
RELEASE_DATE=$(shell date +%FT%T%Z)
PROJECT = github.com/takama/caldera

LDFLAGS = "-s -w \
	-X $(PROJECT)/pkg/version.RELEASE=$(RELEASE) \
	-X $(PROJECT)/pkg/version.DATE=$(RELEASE_DATE)"

GO_PKG = $(shell go list $(PROJECT)/pkg/...)

all: run

vendor: bootstrap
	@echo "+ $@"
	@dep ensure -vendor-only

run: clean build
	@echo "+ $@"
	./${APP}

build: vendor test lint
	@echo "+ $@"
	@go build -a -ldflags $(LDFLAGS) -o $(APP) $(PROJECT)/cmd/caldera

test:
	@echo "+ $@"
	@go test -race -cover $(GO_PKG)

fmt:
	@echo "+ $@"
	@go list -f '"gofmt -w -s -l {{.Dir}}"' $(GO_PKG) | xargs -L 1 sh -c

lint: bootstrap
	@echo "+ $@"
	@golangci-lint run --enable-all ./...

version:
	@./bumper.sh

clean:
	@rm -f ./${APP}

HAS_DEP := $(shell command -v dep;)
HAS_LINT := $(shell command -v golangci-lint;)
HAS_IMPORTS := $(shell command -v goimports;)

bootstrap:
ifndef HAS_DEP
	go get -u github.com/golang/dep/cmd/dep
endif
ifndef HAS_LINT
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
endif
ifndef HAS_IMPORTS
	go get -u golang.org/x/tools/cmd/goimports
endif


.PHONY: all \
	vendor \
	run \
	build \
	test \
	fmt \
	lint \
	version \
	clean \
	bootstrap