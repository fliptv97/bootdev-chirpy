# Chirpy

## Getting started

### Requirements
- [Go](https://go.dev/)
  - [sqlc](https://sqlc.dev/)
  - [goose](https://github.com/pressly/goose)
- [PostgreSQL@15](https://www.postgresql.org/)

### Configuration
First of all, we should apply all migrations from `sql/schema`:
```shell
$ goose postgres "user={user} dbname=chirpy sslmode=disable" up
```

After that we should create and fill out `.env`
```shell
$ cp .env.example .env
```

Now we can start our application from root of the project with simple command
```shell
$ go run .
```

## Troubleshooting
- Be sure, that postgresql service is running. You can check that by running
  ```
  brew services info postgresql@15
  ```
  It should show, that process is loaded and running. If process isn't running, you should start it with:
  ```
  brew services start postgresql@15
  ```