include .env
export

migrate_up_all:
	migrate -path ./migrations -database "postgres://$(PG_USER):$(PG_PASS)@localhost:5432/$(PG_DB)?sslmode=disable" up

migrate_down:
	migrate -path ./migrations -database "postgres://$(PG_USER):$(PG_PASS)@localhost:5432/$(PG_DB)?sslmode=disable" down 

dev_up:
	docker compose -f docker-compose.yml up -d

dev_down:
	docker compose -f docker-compose.yml down

run:
	go run ./cmd/app

testdb_up:
	docker compose -f docker-compose-test.yml up -d

testdb_down:
	docker compose -f docker-compose-test.yml down