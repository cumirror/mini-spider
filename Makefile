# init project path
HOMEDIR := $(shell pwd)
OUTDIR  := $(HOMEDIR)/output
BIN     := mini_spider

VERSION=`git describe --always`
BUILDTIME=`date +%FT%T%z`
GOVERSION=`go version`

# init command params
GO      := go
GOOS    := linux
GOBUILD := $(GO) build
GOTEST  := $(GO) test
GOPKGS  := $$($(GO) list ./pkg... | grep -vE "vendor")

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X 'main.BuildTime=${BUILDTIME}' -X 'main.GoVersion=${GOVERSION}'"

# make, make all
all: prepare compile package

# make prepare
prepare:

# make compile, go build
compile: build
build:
	GOOS=$(GOOS) $(GOBUILD) ${LDFLAGS} -o $(HOMEDIR)/$(BIN) cmd/mini_spider.go

# make test, test your code
test: test-case
test-case:
	$(GOTEST) -v -cover $(GOPKGS)
	rm -r test/webpage

# make package
package: package-bin
package-bin:
	mkdir -p $(OUTDIR)
	mv $(HOMEDIR)/$(BIN) $(OUTDIR)/

# make clean
clean:
	rm -rf $(OUTDIR)

# avoid filename conflict and speed up build
.PHONY: all prepare compile test package clean build
