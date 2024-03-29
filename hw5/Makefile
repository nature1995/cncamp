VERSION = 1.4.0
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
	@go build -o ./bin/$(APP_NAME) ./cmd/main.go

run: dep
	@echo ">> run the application"
	@go run ./cmd/main.go

image:
	@echo ">> build image"
	@docker build -t $(DOCKER_REPO_ACCOUNT)/$(APP_NAME):$(VERSION) -t $(DOCKER_REPO_ACCOUNT)/$(APP_NAME):latest --platform linux/amd64 -f ./Dockerfile .

push-no-build:
	@echo ">> push image to docker hub"
	@docker push $(DOCKER_REPO_ACCOUNT)/$(APP_NAME):$(VERSION)
	@docker push $(DOCKER_REPO_ACCOUNT)/$(APP_NAME):latest

push: image
	@echo ">> push image to docker hub"
	@docker push $(DOCKER_REPO_ACCOUNT)/$(APP_NAME):$(VERSION)
	@docker push $(DOCKER_REPO_ACCOUNT)/$(APP_NAME):latest

deployment:
	@echo ">> deploy to kubernetes"
	@kubectl apply -f ./deploy/k8s/config.yaml
	@kubectl apply -f ./deploy/k8s/deployment.yaml
	@kubectl apply -f ./deploy/k8s/service.yaml
	@kubectl apply -f ./deploy/k8s/ingress.yaml

deployment-hpa:
	@echo ">> deploy hpa to kubernetes"
	@kubectl apply -f ./deploy/k8s/workload.yaml

delete:
	@echo ">> delete from kubernetes"
	@kubectl delete -f ./deploy/k8s/ingress.yaml
	@kubectl delete -f ./deploy/k8s/service.yaml
	@kubectl delete -f ./deploy/k8s/deployment.yaml
	@kubectl delete -f ./deploy/k8s/config.yaml

delete-hpa:
	@echo ">> delete from kubernetes"
	@kubectl delete -f ./deploy/k8s/workload.yaml