createdb:
	docker exec -it postgres-ohKj psql -U postgres -d postgres -c "CREATE DATABASE simple_bank;"
dropdb:
	docker exec -it postgres-ohKj psql -U postgres -d postgres -c "DROP DATABASE simple_bank;"

migrateup:
	migrate -path db/migration -database "postgres://postgres:postgrespw@localhost:49153/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://postgres:postgrespw@localhost:49153/simple_bank?sslmode=disable" -verbose down

makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
sqlc:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

maintest:
	go test -timeout 30s github.com/techshool/simplebank/db/sqlc -run ^TestMain$

test:
	go test -v -cover ./...

.PHONY: createdb dropdb migrateup migratedown sqlc maintest test