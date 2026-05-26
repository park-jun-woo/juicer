# juicer

[![Version](https://img.shields.io/badge/version-v0.1.0-blue.svg)](https://github.com/park-jun-woo/juicer/releases)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![skills.sh](https://skills.sh/b/park-jun-woo/juicer)](https://skills.sh/park-jun-woo/juicer)

**Stop writing API specs by hand.** juicer reads your web framework source code and extracts OpenAPI specs, DDL schemas, and SQL query skeletons — automatically.

- OpenAPI 3.0 spec from source code in seconds, not hours
- DDL migrations merged into clean per-table snapshots
- sqlc query scaffolding with ratchet workflow
- Zero runtime overhead — pure static analysis, no instrumentation

## Quickstart

```bash
npx skills add park-jun-woo/juicer
```

Or install the CLI directly:

```bash
go install github.com/park-jun-woo/juicer/cmd/juicer@latest
```

Then scan your project:

```bash
juicer scan --openapi ./my-gin-project
```

## Supported Frameworks

- **Go + Gin** (current)
- NestJS (planned)
- FastAPI (planned)

## Usage

### Extract OpenAPI 3.0

```bash
juicer scan --openapi ./my-gin-project
juicer scan --openapi -o api.yaml ./my-gin-project
```

### Extract endpoint index (YAML/JSON)

```bash
juicer scan ./my-gin-project
juicer scan -json ./my-gin-project
```

### Parse DDL migrations

```bash
juicer ddl ./migrations -o ./schema
```

### SQL query scaffolding (ratchet workflow)

```bash
juicer sql next --repo ./repository --queries ./db/query
juicer sql status
```

### Hurl API test session

```bash
juicer hurl next --host http://localhost:8080 --tests ./tests --repo ./repository
juicer hurl status
```

## What it extracts

| Layer | Output |
|---|---|
| Routes | HTTP method, path, handler location, middleware |
| Request | Body binding type + struct fields, query/form/path params, file uploads |
| Response | Status codes, body types + struct fields, `json`/`validate` tags |
| OpenAPI | Paths, parameters, requestBody, responses, components/schemas |
| DDL | Per-table CREATE TABLE snapshots from migration history |
| SQL | Repository method skeletons with CRUD type, tables, params, returns |

## Flags

```
juicer scan [flags] [project-root]

  --openapi   Output OpenAPI 3.0 YAML
  --json      Output JSON
  -o string   Write to file instead of stdout

juicer ddl [flags] [migrations-dir]

  -o string   Output directory (one .sql file per table)

juicer sql [flags] [repository-dir]

  --json      Output JSON (default YAML)
  -o string   Output file path
```

## License

MIT
