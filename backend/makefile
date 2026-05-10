MIGRATION_NAME ?= $(name)
ifeq ($(MIGRATION_NAME),)
MIGRATION_NAME := $(word 2,$(MAKECMDGOALS))
endif

ifneq ($(word 2,$(MAKECMDGOALS)),)
$(eval $(word 2,$(MAKECMDGOALS)):;@:)
endif

run:
	go run ./go/cmd/main.go

create-volume:
	podman volume create pg-data-webhooks
start-postgres:
	podman run -d --replace --name webhook-postgres -e POSTGRES_PASSWORD=password -v pg-data-webhooks:/var/lib/postgresql -p 5000:5432 docker.io/library/postgres:latest

migrate-up:
	migrate -path ./db/migrations -database "" up
migrate-down:
	migrate -path ./db/migrations -database "" down

migrate-create:
	migrate create -ext sql -dir ./db/migrations "$(MIGRATION_NAME)"
