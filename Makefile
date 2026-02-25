MIGRATIONS_PATH = ./cmd/migrate/migrations
DB_URL = "postgres://admin:adminpassword@db:5432/orderDB?sslmode=disable"


.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
# migrate-up:
# 	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_URL) up
migrate-up:
	docker compose run --rm migrate
	
.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_URL) down $(filter-out $@,$(MAKECMDGOALS))
