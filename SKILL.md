---
name: codistill
description: Static analysis tool that extracts OpenAPI specs, DDL schemas, and SQL skeletons from 16 web frameworks (Go+Gin, Fiber, Echo, NestJS, Fastify, Hono, FastAPI, Flask, Django, Express, Spring Boot, Quarkus, ASP.NET Core, Supabase Edge Functions, Actix Web, Laravel). Use when you need to generate or update API specs, extract endpoint metadata, manage DDL migrations, scaffold sqlc queries, or merge scan results with existing OpenAPI specs.
license: MIT
metadata:
  author: park-jun-woo
  version: "0.1.11"
---

# codistill — Extract structured specs from web framework source code

## When to Use This Skill

- Extract OpenAPI 3.0 specs from Go+Gin, Fiber, Echo, NestJS, Fastify, Hono, FastAPI, Flask, Django, Express, Spring Boot, Quarkus, ASP.NET Core, Supabase Edge Functions, Actix Web, or Laravel projects
- Merge scan results with an existing openapi.yaml (router registration as ground truth)
- Generate endpoint index (routes, request/response types, middleware) as YAML/JSON
- Parse and merge DDL migration files into per-table snapshots (ALTER COLUMN supported)
- Extract SQL queries from Go repository patterns and scaffold sqlc query files

## When NOT to Use This Skill

- You need runtime request/response capture (use a proxy or instrumentation tool instead)
- The target framework is not one of the 16 supported frameworks listed above

## Install

Go is required. If not installed: https://go.dev/dl/

```bash
git clone https://github.com/park-jun-woo/codistill.git
cd codistill && make install
```

Requires Go 1.25+, CGo enabled, and a C compiler (gcc/clang) for tree-sitter (NestJS/FastAPI support).

## Commands

| Command | Purpose |
|---|---|
| `codist scan [project-root]` | Extract endpoint index as YAML |
| `codist scan --openapi [project-root]` | Generate OpenAPI 3.0 YAML (auto-merges with existing spec) |
| `codist scan --json [project-root]` | Extract endpoint index as JSON |
| `codist scan --framework <fw> [project-root]` | Override framework detection (gogin, fiber, echo, nestjs, fastify, hono, fastapi, flask, django, express, spring, quarkus, dotnet, supafunc, actix, laravel) |
| `codist scan --base <spec> [project-root]` | Explicit base OpenAPI spec for merge |
| `codist ddl [migrations-dir]` | Parse DDL migrations into per-table snapshots |
| `codist sql [repository-dir]` | Extract SQL query skeletons from Go repository code |
| `codist sql next` | Scaffold next sqlc query file (ratchet workflow) |
| `codist sql status` | Show sqlc session progress |
| `codist sql list` | List all queries in session |
| `codist sql skip` | Skip current query |
| `codist sql reset` | Reset sqlc session |

## Workflow

### 1. Scan endpoints

```bash
codist scan ./my-project
codist scan --openapi -o api.yaml ./my-project
```

Framework is auto-detected. Override with `--framework gogin|fiber|echo|nestjs|fastify|hono|fastapi|flask|django|express|spring|quarkus|dotnet|supafunc|actix|laravel`.

### 2. Parse DDL migrations

```bash
codist ddl ./migrations -o ./schema
```

### 3. Extract and scaffold SQL queries

```bash
codist sql ./repository
codist sql next --repo ./repository --queries ./db/query
```

## Key Concepts

- **Multi-framework** — 16 frameworks: Go+Gin, Fiber, Echo (Go), NestJS, Fastify, Hono, Express (TypeScript), FastAPI, Flask, Django (Python), Spring Boot, Quarkus (Java), ASP.NET Core (C#), Supabase Edge Functions (Deno), Actix Web (Rust), Laravel (PHP). Auto-detected from project files.
- **OpenAPI merge** — If existing openapi.yaml found, merges with scan results. Structure from code (ground truth), descriptions from existing spec. Dead specs dropped with warning.
- **Static analysis only** — No runtime, no reflection, no instrumentation.
- **Ratchet workflow** — `sql next` iterates through items one by one. Progress saved in `.codist/` session files.
- **1-depth call tracking** (Go+Gin) — Follows handler wrapper functions that pass `*gin.Context` to recover actual status codes and response types.
- **DDL full lifecycle** — CREATE/DROP TABLE, ADD/DROP COLUMN, ALTER COLUMN (NOT NULL, DEFAULT, TYPE), ADD/DROP CONSTRAINT, CREATE/DROP INDEX.

## Common Errors and Fixes

| Error | Cause | Fix |
|---|---|---|
| `no Go files found` | Wrong project root path | Point to directory containing `go.mod` |
| `no endpoints found` | No framework router detected | Ensure project uses one of the 16 supported frameworks |
| `unknown framework` | Auto-detection failed | Use `--framework gogin\|fiber\|echo\|nestjs\|fastify\|hono\|fastapi\|flask\|django\|express\|spring\|quarkus\|dotnet\|supafunc\|actix\|laravel` |
| `no session found` | Running `next/status/skip` before first session | Run `codist sql next --repo ... --queries ...` first |
| `sqlc generate failed` | sqlc not installed or misconfigured | Install sqlc and ensure `sqlc.yaml` exists |

## Conventions

- Output defaults to YAML on stdout. Use `--json` for JSON, `-o` for file output.
- Session state is stored in `.codist/` directory (gitignored).
- Flags must come before the project path argument.

## Full Documentation

- Source: https://github.com/park-jun-woo/codistill
