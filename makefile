GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
BUILDDIR=$(PWD)/.bin

all: prep test aws-assume-role darwin linux windows
.PHONY: prep
prep:
	mkdir -p $(BUILDDIR)
.PHONY: test
test: 
	$(GOTEST) -v ./...
.PHONY: aws-assume-role 
aws-assume-role:
	$(GOBUILD) -o $(BUILDDIR)/aws-assume-role ./cmd/aws-assume-role
.PHONY: darwin
darwin:
	export GOOS=darwin
	export GOARCH=arm64
	$(GOBUILD) -o $(BUILDDIR)/aws-assume-role_darwin_amd64 ./cmd/aws-assume-role
.PHONY: linux
linux:
	export GOOS=linux 
	export GOARCH=arm64 
	$(GOBUILD) -o $(BUILDDIR)/aws-assume-role_linux_amd64 ./cmd/aws-assume-role
.PHONY: windows
windows:
	export GOOS=windows 
	export GOARCH=arm64 
	$(GOBUILD) -o $(BUILDDIR)/aws-assume-role_windows_amd64.exe ./cmd/aws-assume-role
