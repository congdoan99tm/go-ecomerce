GOOSE_DBSTRING= "root:root1234@tcp(127.0.0.1:33306)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema


# name app
APP_NAME = server

run:
	go run ./cmd/${APP_NAME}/

dev:

kill:
	docker compose kill

up:
	docker compose up -d

down:
	docker compose down

upse:
		@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
		@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

resetse:
		@GOOSE_DRIVER=mysql GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

.PHONY: upse downse resetse

.PHONY: run

.PHONY: air