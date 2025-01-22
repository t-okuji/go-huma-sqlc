# Go Huma sqlc

This project is a learning project for `Go`, `huma` and `sqlc`, and its project structure is based on clean architecture.Using `air` for hot reloading.

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) - Open platform for developing, shipping, and running applications
- [atlas](https://atlasgo.io/) - Migration tool
- [sqlc](https://sqlc.dev/) - Compile SQL to type-safe code

### How To Run

```sh
# Start services
$ docker compose up --build -d

# Database migration
$ atlas schema apply \
--url "postgres://user:password@localhost:5432/demo?sslmode=disable" \
--dev-url "docker://postgres/17-alpine" \
--to file://app/db/schema.sql

# Generate code
$ sqlc generate -f app/sqlc.yaml
```

### Check docs

Open your browser at http://localhost:8080/docs

## Build image

```sh
# Move directory
$ cd app

# Build
$ docker build --target prod -t go-huma-sqlc .

# Run image for Mac (When you want to connect to db on localhost)
$ docker run \
-e POSTGRES_DB=demo \
-e POSTGRES_USER=user \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_PORT=5432 \
-e POSTGRES_HOST=host.docker.internal \
-p 8888:8888 \
go-huma-sqlc

# Run image for linux (When you want to connect to db on localhost)
$ docker run \
-e POSTGRES_DB=demo \
-e POSTGRES_USER=user \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_PORT=5432 \
-e POSTGRES_HOST=localhost \
--net=host \
go-huma-sqlc
```

### Check docs

Open your browser at http://localhost:8888/docs