VERSION = 1.0.0
APP_NAME=http_server

all: dep test build

dep:
	@go mod tidy

test: dep
	@echo ">> running tests"
	@go test -v -gcflags=all=-l ./...

build: dep
	@echo ">> building binaries"
	@go build -o $(APP_NAME) ./main.go

run: dep
	@echo ">> run the application"
	@go run main.go
