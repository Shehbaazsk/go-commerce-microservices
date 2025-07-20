MIGRATE=migrate
MIGRATIONS_DIR=./db/migrations

# Load env vars from .env file
include .env
export 

create_migration:
	@read -p "Migration name: " name; \
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name

migrate_up:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database $(DB_URL) up

migrate_down:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database $(DB_URL) down 1

migrate_drop:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database $(DB_URL) drop -f

migrate_version:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database $(DB_URL) version
