# Define variables
BUILD_DIR = bin/api
include .env
export $(shell sed 's/=.*//' .env)

DB_STRING ?= postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

build:
	@echo "Building Application"
	@go build -o $(BUILD_DIR) cmd/main.go

dev:
	@echo "Starting app in development"
	@go run cmd/main.go

docker:
	@echo "Building docker image"
	@docker build  . -t api:latest

migration-new:
	@echo "Generating new migration file"
	@mkdir -p migrations
	@goose -dir migrations create $(name) sql

migration-up:
	@echo "Running all up migrations"
	 @goose -dir migrations postgres "$(DB_STRING)" up

migration-down:
	@echo "Running all down migrations"
	@goose -dir migrations postgres "$(DB_STRING)" down

help:
	@echo "- dev: 	Start app in development mode"
	@echo "- build: Build app in production mode"
	@echo "- migration-new:  generate new migration file"
	@echo "- migration-up:  run all up migrations"
	@echo "- migration-down:  run all down migrations"
	@echo "- docker: Build docker image"
	@echo "- help:  Show this help message"