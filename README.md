# E-commerce App

A simple e-commerce application built with Go, PostgreSQL, and Docker.

## Prerequisites

- Go 1.25.5 or later
- Docker and Docker Compose
- [sqlc](https://sqlc.dev/) for generating Go code from SQL
- [goose](https://github.com/pressly/goose) for database migrations

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/techies/ecom.git
   cd ecom
   ```

2. **Start the PostgreSQL database:**
   ```bash
   docker compose up -d
   ```
   This will start a PostgreSQL container on port 5442.

3. **Run database migrations:**
   ```bash
   goose up
   ```
   This applies the database schema migrations.

4. **Generate Go code from SQL:**
   ```bash
   sqlc generate
   ```
   This generates the database query code.

5. **Install dependencies:**
   ```bash
   go mod tidy
   ```

6. **Build the application:**
   ```bash
   go build -o ecom cmd/*.go
   ```

7. **Run the application:**
   ```bash
   ./ecom
   ```
   Or directly with:
   ```bash
   go run cmd/*.go
   ```

   The server will start on `http://localhost:8080`.

## API Endpoints

- `GET /health` - Health check
- `GET /products` - List all products

## Testing

Test the API with curl:

```bash
curl http://localhost:8080/health
curl http://localhost:8080/products
```

## Development

- The database connection uses the DSN from the `GOOSE_DBSTRING` environment variable (defined in `.env`).
- Code generation: Run `sqlc generate` after modifying SQL queries in `internal/adapters/postgres/sqlc/`.
- Migrations: Add new migration files in `internal/adapters/postgres/migrations/` and run `goose up`.

## Stopping

To stop the database:
```bash
docker compose down
```