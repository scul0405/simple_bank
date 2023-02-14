# SIMPLE BANK

This repository contains a project I have learn about design, develop and deploy a complete backend system using Go, PostgreSQL and Docker.

## Setup local development

#### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [DB Docs](https://dbdocs.io/docs)
- [DBML CLI](https://www.dbml.org/cli/#installation)
- [Sqlc](https://github.com/kyleconroy/sqlc#installation)
- [Gomock](https://github.com/golang/mock)

## Quick start

Clone this repository:
```sh
git clone https://github.com/scul0405/simple_bank.git
cd simple_bank
```
#### Using with docker compose

Run this command:
```sh
docker compose up
```
#### Using by default

On Linux/Mac use `make` to execute Makefile command (if you use Windows, you can use `mingw32-make` instead or rename `mingw32-make` to `make`):

Install and run PostgreSQL docker container:
```sh
    make postgres
```

Create database:
```sh
    make createdb
```

Migrate database:
```sh
    make migrateup
```

Run server:
```sh
    make server
```

Now server is started on [http://localhost:3000/](http://localhost:3000/) for HTTP server and [http://localhost:8080/](http://localhost:8080/) for gRPC server. You can use Evans to send request to gRPC server by use this command:
```sh
    make evans
```

## API endpoints
#### With gateway server
* `POST /v1/create_user`: Create a new user
* `POST /v1/login_user`: Login user and get access token & refresh token

#### With Gin server
You can change to Gin server by change `line:39, 40` in `main.go` to `runGinServer(config, store)`

* `POST /users`: Create a new user
* `POST /users/login`: Login user and get access token & refresh token
* `POST /tokens/renew_access`: Return new access token
* `POST /accounts`: Create a new account
* `GET /accounts/:id`: Get an account with account's ID
* `POST /transfers`: Transfer money between 2 accounts

## Testing
Run testing for all packages with random data
```sh
make test
```

## Setup infrastructure

- Start postgres container:
    ```bash
    make postgres
    ```

- Create simple_bank database:
    ```bash
    make createdb
    ```

- Run db migration up all versions:
    ```bash
    make migrateup
    ```

- Run db migration up 1 version:
    ```bash
    make migrateuplast
    ```

- Run db migration down all versions:
    ```bash
    make migratedown
    ```

- Run db migration down 1 version:
    ```bash
    make migratedownlast
    ```

## Documentation

#### Database
- Generate DB documentation:
    ```bash
    make db_docs
    ```

- Access the DB documentation at [this address](https://dbdocs.io/vldtruong1221/Simple_bank). Password: `secret`

#### Swagger (API Documentation)
- Run server:
    ```bash
    make server
    ```

- Go to [http://localhost:3000/swagger](http://localhost:3000/swagger) to view API documentation

## How to generate code

- Generate schema SQL file with DBML:
    ```bash
    make db_schema
    ```

- Generate SQL CRUD with sqlc:
    ```bash
    make sqlc
    ```

- Generate DB mock with gomock:
    ```bash
    make mock
    ```

- Create a new db migration:
    ```bash
    migrate create -ext sql -dir db/migration -seq <migration_name>
    ```
