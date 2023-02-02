postgres:
	docker run --name postgres-ohKj -p 49153:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgrespw -d postgres:latest

createdb:
	docker exec -it postgres-ohKj psql -U postgres -d postgres -c "CREATE DATABASE simple_bank;"
dropdb:
	docker exec -it postgres-ohKj psql -U postgres -d postgres -c "DROP DATABASE simple_bank;"

migrateup:
	migrate -path db/migration -database "postgres://postgres:postgrespw@localhost:49153/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://postgres:postgrespw@localhost:49153/simple_bank?sslmode=disable" -verbose down

migrateuplast:
	migrate -path db/migration -database "postgres://postgres:postgrespw@localhost:49153/simple_bank?sslmode=disable" -verbose up 1
migratedownlast:
	migrate -path db/migration -database "postgres://postgres:postgrespw@localhost:49153/simple_bank?sslmode=disable" -verbose down 1

makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
sqlc:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

maintest:
	go test -timeout 30s github.com/techschool/simplebank/db/sqlc -run ^TestMain$

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc maintest test server mock migrateuplast migratedownlast