#!make
# You can change file env like .env or .env.*.local
include .env

ENGINE=cmd/server/main.go
BUILD_DIR=build
CONN_STRING="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL}"
RPC_PORT=4200

debug:
	go run ${ENGINE} service --svport ${RPC_PORT}
.PHONY: debug

build:
	@echo "Building app"
	go build -o ${BUILD_DIR}/app ${ENGINE}
	@echo "Success build app. Your app is ready to use in 'build/' directory."
.PHONY: build