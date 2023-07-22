.PHONY: default run build db-up db-down db-stop db-start clean

# Variables
APP_NAME=db-go

# Tasks
default: run

run:
	@go run main.go

db-up:
	@docker run -d --name db-golang -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=db-golang -p 3306:3306 mysql:latest

db-down:
	@docker rm db-golang

db-start:
	@docker run db-golang

db-stop:
	@docker stop db-golang

build:
	go build -o ${APP_NAME} main.go

clean:
	@rm -f ${APP_NAME}