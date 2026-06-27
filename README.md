# 🗂️ GoVault

> Your personal notes and tasks — simple, fast, yours.

Welcome to **GoVault** — a lightweight personal manager for notes and tasks, built entirely in Go. No heavy frameworks, no magic. Just clean, idiomatic Go that you can read, understand, and be proud of.

This project was born as a learning journey from C# to Go — and it turned out to be a pretty great one. If you're on the same path, you're in the right place. Pull up a chair.

---

## ✨ What GoVault does

- 📝 Create, read, update and delete notes and tasks
- 🏷️ Tag and filter entries any way you like
- 🌐 Exposes a clean JSON REST API — built with stdlib `net/http`, no Gin, no Fiber
- 💻 Comes with a CLI client so you can use it right from your terminal
- 📤 Export your notes to Markdown in one command

---

## 🛠️ Tech stack

| Layer | Choice | Why |
|---|---|---|
| Language | Go 1.22+ | Obviously |
| Database | SQLite (`mattn/go-sqlite3`) | Zero setup, perfect for local tools |
| Router | stdlib `net/http` + `ServeMux` | Learn the foundation first |
| Config | `os.Getenv` + `.env` file | Simple and explicit |
| Testing | `testing` + `testify` | Idiomatic and widely used |

---

## 📁 Project structure

```
govault/
├── cmd/
│   ├── server/        # Entrypoint: REST API server
│   └── cli/           # Entrypoint: CLI tool
├── internal/
│   ├── store/         # DB layer — interfaces + SQLite implementation
│   ├── handler/       # HTTP handlers
│   └── model/         # Domain structs (Note, Task, Tag...)
├── migrations/        # SQL migration files
├── .env.example
└── README.md
```

The `internal/` directory keeps your business logic private to this module — a Go convention that replaces the access modifiers you know from C#.

---

## 🚀 Getting started

```bash
# Clone the repo
git clone https://github.com/you/govault.git
cd govault

# Run the API server
go run ./cmd/server

# Or use the CLI
go run ./cmd/cli note add "Buy milk"
go run ./cmd/cli task list --tag shopping
```

That's it. No `dotnet restore`, no solution files, no project references. Just Go.

---

## 🧠 What you'll learn building this

If you're coming from C#, GoVault is designed to surface the most important Go concepts naturally:

- **Error handling** — `if err != nil` everywhere, and why that's actually a good thing
- **Interfaces** — `type NoteStore interface { ... }` instead of DI containers and abstract classes
- **Packages over classes** — `internal/store` is your data layer, not a class hierarchy
- **Goroutines** — a background worker for task reminders, using `context` and `os.Signal`
- **Graceful shutdown** — the Go way, with channels and `sync.WaitGroup`
- **No ORM** — raw `database/sql` so you actually understand what's happening

---

## 🔭 Future improvements

GoVault is designed to grow with you. Here's a roadmap ordered by complexity — start from the top and work your way down as you get more comfortable.

### Level 1 — Go deeper
- **Background worker** — remind yourself about due tasks using goroutines and tickers
- **Middleware chain** — add logging and rate limiting via `http.Handler` composition
- **Unit tests with mocks** — swap the SQLite store for an in-memory mock using interfaces
- **Structured logging** — replace `fmt.Println` with `log/slog` (introduced in Go 1.21)

### Level 2 — Production-ready
- **PostgreSQL** — migrate from SQLite using `pgx`, learn connection pooling
- **JWT authentication** — implement it yourself with `crypto/hmac`, no libraries needed
- **Dockerize** — write a multi-stage `Dockerfile` and a `docker-compose.yml`
- **OpenAPI spec** — document your API and generate a client from it

### Level 3 — Advanced Go
- **gRPC endpoint** — add a gRPC interface alongside REST and feel the difference
- **WebSocket** — live task updates using `golang.org/x/net/websocket`
- **Benchmarks** — write `testing.B` benchmarks and profile with `pprof`
- **Plugin system** — load exporters (Markdown, JSON, CSV) as Go plugins at runtime

---

## 🤝 Contributing

This is a personal learning project, but PRs and issues are always welcome. If you spot something un-idiomatic or have a better Go pattern — please share it. That's the whole point.

---

## 📄 License

MIT — do whatever you want with it.
