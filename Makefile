SHELL=/usr/bin/env bash
VERSION  := ${strip ${shell cat VERSION}}

LDFLAGS   := -w -s
BINDIR    := $(CURDIR)/bin
GIT_COMMIT = $(shell git rev-parse HEAD)

LDFLAGS += -X k8s.io/helm/pkg/version.Version=${VERSION}
LDFLAGS += -X k8s.io/helm/pkg/version.GitCommit=${GIT_COMMIT}

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(BINDIR)/tiller -ldflags '$(LDFLAGS)' k8s.io/helm/cmd/tiller

.PHONY: clean
clean:
	@rm -rf $(BINDIR) ./rootfs/tiller ./_dist

.PHONY: bootstrap
bootstrap:
	go build -o bin/protoc-gen-go ./vendor/github.com/golang/protobuf/protoc-gen-go
