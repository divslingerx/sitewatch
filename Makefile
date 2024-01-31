.PHONY: build test clean

# The name of your binary
BINARY_NAME=sitewatcher

# The Go path
GOPATH=$(shell go env GOPATH)

# The Go build command
GOBUILD=go build

# The Go test command
GOTEST=go test -v

# The Go clean command
GOCLEAN=go clean

build:
	cd cmd/sitewatcher && $(GOBUILD) -o bin/$(BINARY_NAME) -v

test:
	$(GOTEST) ./...

clean:
	$(GOCLEAN)
	rm -f $(GOPATH)/bin/$(BINARY_NAME)

run:
	$(GOPATH)/bin/$(BINARY_NAME)
