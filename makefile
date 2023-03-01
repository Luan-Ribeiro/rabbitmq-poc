SHELL := /bin/bash

# =====================================
# Variable

VERSION := 1.0
PROJECT_NAME := kafka-poc

run/:
	go run ./producer/main.go
	go run ./consumer/main.go

## Download libraries
tidy:
	go mod tidy
	go mod vendor

# =====================================
# Enviroment

env-up:
	docker-compose up -d

env-down:
	docker-compose down