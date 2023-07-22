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
	@docker stop db-golang
	@docker rm db-golang

db-start:
	@docker run db-golang

db-stop:
	@docker stop db-golang

migration:
	@docker exec -i db-golang mysql -uroot -proot db-golang < ./migrations/init.sql

db-populate:
	@docker exec -i db-golang mysql -uroot -proot db-golang < ./migrations/populate.sql

build:
	go build -o ${APP_NAME} main.go

clean:
	@rm -f ${APP_NAME}