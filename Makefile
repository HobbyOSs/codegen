# Go parameters
GOBUILD=go build
GOTEST=gotest

BIN=codegen

.PHONY: all test

all: dep build test

build: dep gen
	$(GOBUILD) -v
	go install -v ./...

test:
	go install -v github.com/rakyll/gotest@latest
	$(GOTEST) -v ./...

clean:
	go clean

run: build
	chmod +x ./$(BIN)
	./$(BIN)

gen:
	go generate ./...

fmt:
	go fmt ./...

dep:
	go mod download
	go mod tidy

tool:
	go install -v golang.org/x/tools/gopls@latest
	go install -v github.com/go-delve/delve/cmd/dlv@latest
	go install github.com/dmarkham/enumer@latest
