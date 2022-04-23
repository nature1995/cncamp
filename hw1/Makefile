VERSION = 1.0.0
SERVER_BIN = "hw1"

all: test build

build:
	@echo ">> building binaries"
	@go build -o $(SERVER_BIN) ./main.go

test:
	@echo ">> running tests"
	@go test -v -gcflags=all=-l ./...
