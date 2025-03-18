# Define variables
BUILD_DIR = bin/api

build:
	@echo "Building Application"
	@go build -o $(BUILD_DIR) cmd/main.go

dev:
	@echo "Starting app in development"
	@go run cmd/main.go

docker:
	@echo "Building docker image"
	@docker build  . -t api:latest

help:
	@echo "- dev: 	Start app in development mode"
	@echo "- build: Build app in production mode"
	@echo "- docker: Build docker image"