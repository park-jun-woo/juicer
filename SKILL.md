---
name: juicer
description: Static analysis tool that extracts OpenAPI specs, DDL schemas, and SQL skeletons from web framework projects (Go+Gin, NestJS, FastAPI). Use when you need to generate or update API specs, extract endpoint metadata, manage DDL migrations, scaffold sqlc queries, or merge scan results with existing OpenAPI specs.
license: MIT
metadata:
  author: park-jun-woo
  version: "0.1.0"
---

# juicer — Extract structured specs from web framework source code

## When to Use This Skill

- Extract OpenAPI 3.0 specs from Go+Gin, NestJS, or FastAPI projects
- Merge scan results with an existing openapi.yaml (router registration as ground truth)
- Generate endpoint index (routes, request/response types, middleware) as YAML/JSON
- Parse and merge DDL migration files into per-table snapshots (ALTER COLUMN supported)
- Extract SQL queries from Go repository patterns and scaffold sqlc query files

## When NOT to Use This Skill

- You need runtime request/response capture (use a proxy or instrumentation tool instead)
- The target framework is not Go+Gin, NestJS, or FastAPI

## Install

Go is required. If not installed: https://go.dev/dl/

```bash
git clone https://github.com/park-jun-woo/juicer.git
cd juicer && make install
```

Requires Go 1.25+, CGo enabled, and a C compiler (gcc/clang) for tree-sitter (NestJS/FastAPI support).

## Commands

| Command | Purpose |
|---|---|
| `juicer scan [project-root]` | Extract endpoint index as YAML |
| `juicer scan --openapi [project-root]` | Generate OpenAPI 3.0 YAML (auto-merges with existing spec) |
| `juicer scan --json [project-root]` | Extract endpoint index as JSON |
| `juicer scan --framework <fw> [project-root]` | Override framework detection (gogin, nestjs, fastapi) |
| `juicer scan --base <spec> [project-root]` | Explicit base OpenAPI spec for merge |
| `juicer ddl [migrations-dir]` | Parse DDL migrations into per-table snapshots |
| `juicer sql [repository-dir]` | Extract SQL query skeletons from Go repository code |
| `juicer sql next` | Scaffold next sqlc query file (ratchet workflow) |
| `juicer sql status` | Show sqlc session progress |
| `juicer sql list` | List all queries in session |
| `juicer sql skip` | Skip current query |
| `juicer sql reset` | Reset sqlc session |

## Workflow

### 1. Scan endpoints

```bash
juicer scan ./my-project
juicer scan --openapi -o api.yaml ./my-project
```

Framework is auto-detected. Override with `--framework gogin|nestjs|fastapi`.

### 2. Parse DDL migrations

```bash
juicer ddl ./migrations -o ./schema
```

### 3. Extract and scaffold SQL queries

```bash
juicer sql ./repository
juicer sql next --repo ./repository --queries ./db/query
```

## Key Concepts

- **Multi-framework** — Go+Gin (`go/ast` + `go/types`), NestJS (tree-sitter TypeScript), FastAPI (tree-sitter Python). Auto-detected from project files.
- **OpenAPI merge** — If existing openapi.yaml found, merges with scan results. Structure from code (ground truth), descriptions from existing spec. Dead specs dropped with warning.
- **Static analysis only** — No runtime, no reflection, no instrumentation.
- **Ratchet workflow** — `sql next` iterates through items one by one. Progress saved in `.juicer/` session files.
- **1-depth call tracking** (Go+Gin) — Follows handler wrapper functions that pass `*gin.Context` to recover actual status codes and response types.
- **DDL full lifecycle** — CREATE/DROP TABLE, ADD/DROP COLUMN, ALTER COLUMN (NOT NULL, DEFAULT, TYPE), ADD/DROP CONSTRAINT, CREATE/DROP INDEX.

## Common Errors and Fixes

| Error | Cause | Fix |
|---|---|---|
| `no Go files found` | Wrong project root path | Point to directory containing `go.mod` |
| `no endpoints found` | No framework router detected | Ensure project uses Gin/NestJS/FastAPI |
| `unknown framework` | Auto-detection failed | Use `--framework gogin\|nestjs\|fastapi` |
| `no session found` | Running `next/status/skip` before first session | Run `juicer sql next --repo ... --queries ...` first |
| `sqlc generate failed` | sqlc not installed or misconfigured | Install sqlc and ensure `sqlc.yaml` exists |

## Conventions

- Output defaults to YAML on stdout. Use `--json` for JSON, `-o` for file output.
- Session state is stored in `.juicer/` directory (gitignored).
- Flags must come before the project path argument.

## Full Documentation

- Source: https://github.com/park-jun-woo/juicer
