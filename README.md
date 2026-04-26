# Running the application

This document explains how to prepare the SQLite database (run migrations) and start the GraphQL server included in this repository.

Prerequisites
- Go (1.17+)
- sqlite3 CLI (for running the SQL migration files)
- (Optional) Any GraphQL client or just `curl` to test the endpoint

1) Configure environment
The application reads the database DSN from the `DATABASE_URI` environment variable (the project uses `github.com/joho/godotenv` to load `.env` files). Create a `.env` file in the project root with something like:

```
DATABASE_URI=./dev.db
```

The value above will create/use a SQLite database file named `dev.db` in the project root. You can use a different path if you prefer.

2) Run database migrations (SQLite)
This repository contains SQL migration files under `database/migrations/`. Apply the "up" migrations in order so the schema is created and sample data is inserted.

Using the `sqlite3` CLI you can apply each migration file by redirecting it into the database file. From the project root:

- Create the database file and apply the schema migration:
```
sqlite3 ./dev.db < database/migrations/000001_cars_table.up.sql
```

- (Optional) If there are additional migrations that insert seed data, run them next in chronological order:
```
sqlite3 ./dev.db < database/migrations/20260424134429_insert_cars.up.sql
```

Notes:
- The migration files are ordered by their filename prefix. Apply them in increasing order if there are multiple.
- If you prefer, you can open an interactive sqlite3 session and use the `.read` command:
```
sqlite3 ./dev.db
sqlite> .read database/migrations/000001_cars_table.up.sql
sqlite> .read database/migrations/20260424134429_insert_cars.up.sql
sqlite> .exit
```

If you use a migration tool (for example `golang-migrate`), consult that tool's documentation for the correct driver DSN format and run the migrations from the `database/migrations` folder. The SQL files here are standard SQL and should work with any tool that runs raw SQL migration files against a SQLite database.

3) Start the server
From the project root you can run the server directly:

```
go run main.go
```

Or build and run the binary:

```
go build -o bin/graphql .
./bin/graphql
```

By default the server in `main.go` listens on `127.0.0.1:5000` and registers a handler at `/graphql`.

4) Test the GraphQL endpoint
You can query the server using `curl` (Linux / macOS style single-quotes). Example: list all cars:

```
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"query":"query { cars { id brand model year price } }"}' \
  http://127.0.0.1:5000/graphql
```

On Windows PowerShell you may need to adjust quoting, for example:

```
curl -X POST -H "Content-Type: application/json" -d "{\"query\":\"query { cars { id brand model year price } }\"}" http://127.0.0.1:5000/graphql
```

Example response (abbreviated):

```
{
  "data": {
    "cars": [
      { "id": 1, "brand": "Toyota", "model": "Corolla", "year": 2018, "price": 15000.0 },
      { "id": 2, "brand": "Honda", "model": "Civic", "year": 2019, "price": 18000.0 }
    ]
  }
}
```

5) Notes and troubleshooting
- If your database file is empty (no `cars` table), confirm you applied the schema migration (`000001_cars_table.up.sql`).
- The repository also contains a warmup function that can insert sample rows programmatically; the migration files already include a seed migration that inserts sample cars. The Warmup function is available in code if you prefer to seed data programmatically.
- If you see errors connecting to the DB, check `DATABASE_URI` and file permissions.
- If you want to run a GraphQL playground during development, add or enable the playground handler in the server code (this repository may or may not include one by default).
