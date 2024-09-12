
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=dbname=praktikum host=localhost port=5432 user=postgres password=postgres sslmode=disable
export GOOSE_MIGRATION_DIR=./migrations

migrate:
	goose up

migrate_down:
	goose down

create_migration:
	@read -p "Введите название миграции: " migration; goose create $$migration sql