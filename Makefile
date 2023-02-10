DB_URL=postgres://postgres:postgrespw@localhost:5678/simple_bank?sslmode=disable

postgres:
	docker run --name postgres -p 5678:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgrespw -d postgres:latest

createdb:
	docker exec -it postgres psql -U postgres -d postgres -c "CREATE DATABASE simple_bank;"
dropdb:
	docker exec -it postgres psql -U postgres -d postgres -c "DROP DATABASE simple_bank;"

migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up
migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

migrateuplast:
	migrate -path db/migration -database "${DB_URL}" -verbose up 1
migratedownlast:
	migrate -path db/migration -database "${DB_URL}" -verbose down 1

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

db_docs:
	dbdocs build ./doc/db.dbml
db_schema:
	dbml2sql ./doc/db.dbml --postgres -o ./doc/schema.sql

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    proto/*.proto

evans:
	evans --host localhost --port 8080 -r repl

.PHONY: postgres createdb dropdb migrateup migratedown sqlc maintest test server mock migrateuplast migratedownlast db_docs db_schema proto