# Repository Guidelines

## Project Structure & Module Organization

This Go 1.27.3 REST API uses Gin, Cobra, GORM, and Redis. `main.go` starts the CLI; `app/cmd/` contains commands such as `serve`, `migrate`, and `seed`. Routes are in `routes/`; controllers, middleware, requests, models, and policies are under `app/`. Reusable infrastructure is in `pkg/`, startup wiring in `bootstrap/`, defaults in `config/`, and migrations, factories, and seeders in `database/`. Runtime uploads and logs belong in `public/` and `storage/`; do not commit generated data.

## Build, Test, and Development Commands

- `go build ./...` compiles all packages.
- `go test ./...` runs package tests (the repository currently has no checked-in Go test files).
- `cp .env.example .env` creates local configuration; set database and Redis credentials before starting services.
- `go run main.go serve` starts the API server using `.env`.
- `go run main.go migrate up` applies pending database migrations; `go run main.go seed` inserts sample data.
- `go run main.go -h` lists CLI commands and options. Use `-e testing` to load `.env.testing` when present.

## Coding Style & Naming Conventions

Format Go changes with `gofmt` and keep package names short, lowercase, and without underscores. Use exported `PascalCase` names and unexported `camelCase` names. Follow surrounding conventions: feature-specific model and controller code stays in its existing subpackage, and migration filenames use a timestamp plus a descriptive snake_case action (for example, `2022_08_18_111518_add_users_table.go`). Keep configuration defaults in `config/` and environment-specific secrets in ignored `.env` files.

## Testing Guidelines

Add tests alongside the package they cover in `*_test.go` files, using Go's standard `testing` package. Name test functions `Test...` and cover behavior and error cases. Run focused tests with `go test ./path/to/package` and `go test ./...` before submitting. Tests that touch MySQL or Redis should document required services or use isolated configuration.

## Commit & Pull Request Guidelines

Recent commits use short imperative summaries, sometimes in English and sometimes Chinese (for example, `fix env problem` and `缓存友情链接列表`); keep subjects concise and focused on one change. Pull requests should describe behavior changes, list commands or tests run, link related issues, and include request/response examples for API changes. Call out configuration and migration changes.

## Security & Configuration Tips

Never commit credentials, tokens, or production data. Use `.env.example` for safe defaults and update it when adding keys. Review migration effects and avoid destructive operations against shared databases.
