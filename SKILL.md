---
name: juicer
description: Static analysis tool that extracts OpenAPI specs, DDL schemas, and SQL skeletons from web framework projects (Go+Gin, NestJS, FastAPI). Use when you need to generate or update API specs, extract endpoint metadata, manage DDL migrations, scaffold sqlc queries, or run hurl-based API tests from source code.
license: MIT
metadata:
  author: park-jun-woo
  version: "0.1.0"
---

# juicer — Extract structured specs from web framework source code

## When to Use This Skill

- Extract OpenAPI 3.0 specs from a Go+Gin project without runtime or reflection
- Generate endpoint index (routes, request/response types, middleware) as YAML/JSON
- Parse and merge DDL migration files into per-table snapshots
- Extract SQL queries from Go repository patterns and scaffold sqlc query files
- Run hurl-based API test sessions with ratchet workflow (next/skip/pass/fail)

## When NOT to Use This Skill

- The project uses oapi-codegen or similar generators that already produce OpenAPI specs
- You need runtime request/response capture (use a proxy or instrumentation tool instead)
- The target framework is not yet supported (currently Go+Gin only)

## Install

```bash
go install github.com/park-jun-woo/juicer@latest
```

## Commands

| Command | Purpose |
|---|---|
| `juicer scan [project-root]` | Extract endpoint index as YAML |
| `juicer scan --openapi [project-root]` | Generate OpenAPI 3.0 YAML |
| `juicer scan --json [project-root]` | Extract endpoint index as JSON |
| `juicer ddl [migrations-dir]` | Parse DDL migrations into per-table snapshots |
| `juicer sql [repository-dir]` | Extract SQL query skeletons from Go repository code |
| `juicer sql next` | Scaffold next sqlc query file (ratchet workflow) |
| `juicer sql status` | Show sqlc session progress |
| `juicer sql list` | List all queries in session |
| `juicer sql skip` | Skip current query |
| `juicer sql reset` | Reset sqlc session |
| `juicer hurl next` | Run next hurl test (ratchet workflow) |
| `juicer hurl status` | Show hurl session progress |
| `juicer hurl list` | List all endpoints in session |
| `juicer hurl skip` | Skip current endpoint |
| `juicer hurl reset` | Reset hurl session |

## Workflow

### 1. Scan endpoints

```bash
juicer scan ./my-gin-project
juicer scan --openapi -o api.yaml ./my-gin-project
```

### 2. Parse DDL migrations

```bash
juicer ddl ./migrations -o ./schema
```

### 3. Extract and scaffold SQL queries

```bash
juicer sql ./repository
juicer sql next --repo ./repository --queries ./db/query
```

### 4. Run hurl API tests

```bash
juicer hurl next --host http://localhost:8080 --tests ./tests --repo ./repository
juicer hurl status
```

## Key Concepts

- **Static analysis only** — Uses `go/ast` and `go/types`. No runtime, no reflection, no instrumentation.
- **Ratchet workflow** — `sql next` and `hurl next` iterate through items one by one. Each item is todo/done/skipped. Progress is saved in `.juicer/` session files.
- **1-depth call tracking** — Follows handler wrapper functions that pass `*gin.Context` to recover actual status codes and response types.
- **`gin.H` partial confidence** — `map[string]any` responses have keys extracted but value types are best-effort. Marked with `x-schema-confidence: partial`.

## Common Errors and Fixes

| Error | Cause | Fix |
|---|---|---|
| `no Go files found` | Wrong project root path | Point to directory containing `go.mod` |
| `no endpoints found` | No Gin router detected | Ensure project uses `gin.Default()` or `gin.New()` |
| `no session found` | Running `next/status/skip` before first session | Run `juicer sql next --repo ... --queries ...` first |
| `sqlc generate failed` | sqlc not installed or misconfigured | Install sqlc and ensure `sqlc.yaml` exists |

## Conventions

- Output defaults to YAML on stdout. Use `--json` for JSON, `-o` for file output.
- Session state is stored in `.juicer/` directory (gitignored).
- Flags must come before the project path argument.

## Full Documentation

- Source: https://github.com/park-jun-woo/juicer
