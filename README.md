# codistill

<p align="center">
  <img src="codistill.webp" alt="codistill — extract structured specs from web framework source code" width="480">
</p>

[![Version](https://img.shields.io/badge/version-v0.1.1-blue.svg)](https://github.com/park-jun-woo/codistill/releases)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![skills.sh](https://skills.sh/b/park-jun-woo/codistill)](https://skills.sh/park-jun-woo/codistill)

**Stop writing API specs by hand.** codist reads your web framework source code and extracts OpenAPI specs, DDL schemas, and SQL query skeletons — automatically.

- OpenAPI 3.0 spec from source code in seconds, not hours
- Merge with existing openapi.yaml — router registration is ground truth
- DDL migrations merged into clean per-table snapshots (ALTER COLUMN supported)
- sqlc query scaffolding with ratchet workflow
- Zero runtime overhead — pure static analysis, no instrumentation

## Quickstart

```bash
npx skills add park-jun-woo/codistill
```

Or install the CLI directly (requires [Go](https://go.dev/dl/)):

```bash
git clone https://github.com/park-jun-woo/codistill.git
cd codistill && make install
```

Then scan your project:

```bash
codist scan --openapi ./my-project
```

## Supported Frameworks

| Framework | Language | Status |
|---|---|---|
| **Go + Gin** | Go | Stable — `go/ast` + `go/types`, oapi-codegen `.gen.go` supported |
| **NestJS** | TypeScript | Stable — tree-sitter, decorator-based extraction |
| **FastAPI** | Python | Stable — tree-sitter, Pydantic model extraction |

Framework is auto-detected from `go.mod`, `package.json`, or `requirements.txt`. Override with `--framework`:

```bash
codist scan --framework gogin ./project
codist scan --framework nestjs ./project
codist scan --framework fastapi ./project
```

## Usage

### Extract OpenAPI 3.0

```bash
codist scan --openapi ./my-project
codist scan --openapi -o api.yaml ./my-project
```

If the project already has an `openapi.yaml`, codist auto-detects and merges — structure from code, descriptions from existing spec. Dead specs (not registered in router) are dropped.

```bash
# Explicit base spec
codist scan --openapi --base existing-openapi.yaml ./my-project
```

### Extract endpoint index (YAML/JSON)

```bash
codist scan ./my-project
codist scan --json ./my-project
```

### Parse DDL migrations

```bash
codist ddl ./migrations -o ./schema
```

Supports CREATE/DROP TABLE, ADD/DROP COLUMN, ALTER COLUMN (SET/DROP NOT NULL, SET/DROP DEFAULT, TYPE), ADD/DROP CONSTRAINT, CREATE/DROP INDEX.

### SQL query scaffolding (ratchet workflow)

```bash
codist sql next --repo ./repository --queries ./db/query
codist sql status
```

## What it extracts

| Layer | Output |
|---|---|
| Routes | HTTP method, path, handler location, middleware |
| Request | Body binding type + struct fields, query/form/path params, file uploads |
| Response | Status codes, body types + struct fields, `json`/`validate` tags |
| OpenAPI | Paths, parameters, requestBody, responses, components/schemas |
| DDL | Per-table CREATE TABLE snapshots from migration history (ALTER COLUMN supported) |
| SQL | Repository method skeletons with CRUD type, tables, params, returns |

## Flags

```
codist scan [flags] [project-root]

  --openapi       Output OpenAPI 3.0 YAML
  --json          Output JSON
  --framework     Framework override (gogin, nestjs, fastapi)
  --base string   Base OpenAPI spec to merge with
  -o string       Write to file instead of stdout

codist ddl [flags] [migrations-dir]

  -o string   Output directory (one .sql file per table)

codist sql [flags] [repository-dir]

  --json      Output JSON (default YAML)
  -o string   Output file path
```

## Changelog

### v0.1.1

- OpenAPI `$ref` schema generation for named types (placeholder when fields unavailable)
- Unique `operationId` for inherited controllers (path prefix dedup)
- `required`/`enum`/`minLength`/`maxLength`/`minimum`/`maximum` constraints in OpenAPI
- `default: None` → `null` conversion for FastAPI query params
- NestJS generic type substitution in BaseController factory pattern
- OpenAPI `securitySchemes` + per-endpoint `security` from Guard/Depends
- `OmitType([...] as const)` array extraction fix
- `setGlobalPrefix` detection in non-main.ts files
- FastAPI `include_router(module.router)` attribute childVar support
- Primitive type inline schema (`bool` → `type: boolean`, not `$ref`)
- Pydantic/SQLModel inheritance field merging (parent → child)
- Multi-hop `include_router` prefix propagation with convergence loop
- DTO factory (`OmitType`/`PartialType`) optional/validate preservation
- `@Param()` keyless pattern → path template variable name extraction
- `@IsEnum(TaskStatus)` → cross-file enum member value extraction
- FastAPI `required` array from `hasDefault` + `Field(default=...)` analysis
- NestJS barrel `export * from` re-export DTO resolution
- Comma-separated `from ... import a, b, c` dotted_name parsing fix

### v0.1.0

- Initial release
- Go+Gin, NestJS, FastAPI endpoint extraction
- OpenAPI 3.0 spec generation with base spec merging
- DDL migration parsing (CREATE/ALTER/DROP TABLE)
- SQL query scaffolding with ratchet workflow

## License

MIT
