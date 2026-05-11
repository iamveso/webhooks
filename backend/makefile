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
delete-volume:
	podman volume rm pg-data-webhooks
start-postgres:
	podman run -d --replace --name webhook-postgres -e POSTGRES_PASSWORD=password -v pg-data-webhooks:/var/lib/postgresql -p 5000:5432 docker.io/library/postgres:latest

drop-postgres:
	podman stop webhook-postgres; podman rm webhook-postgres

migrate-up:
	migrate -path ./db/migrations -database "postgres://postgres:password@localhost:5000/postgres?sslmode=disable" up
migrate-down:
	migrate -path ./db/migrations -database "postgres://postgres:password@localhost:5000/postgres?sslmode=disable" down

migrate-create:
	migrate create -ext sql -dir ./db/migrations "$(MIGRATION_NAME)"
