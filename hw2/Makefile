VERSION = 1.0.0
APP_NAME=http_server
DOCKER_REPO_ACCOUNT=naturegong

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

image:
	@echo ">> build image"
	@docker build -t $(DOCKER_REPO_ACCOUNT)/$(APP_NAME):$(VERSION) --platform linux/amd64 -f ./Dockerfile .

push: image
	@echo ">> push image to docker hub"
	@docker push $(DOCKER_REPO_ACCOUNT)/$(APP_NAME):$(VERSION)
