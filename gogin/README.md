# huma

Static analysis tool that extracts OpenAPI specs, DDL schemas, and SQL queries from Go+Gin projects.

Feed it a Gin codebase, get structured specs out.

## Install

```bash
go install github.com/park-jun-woo/huma@latest
```

Or build from source:

```bash
git clone https://github.com/park-jun-woo/huma.git
cd huma
make install
```

## Usage

### Extract OpenAPI 3.0

```bash
huma scan --openapi ./my-gin-project
huma scan --openapi -o api.yaml ./my-gin-project
```

### Extract endpoint index (YAML/JSON)

```bash
huma scan ./my-gin-project
huma scan -json ./my-gin-project
```

## What it extracts

| Layer | Output |
|---|---|
| Routes | HTTP method, path, handler location, middleware |
| Request | Body binding type + struct fields, query/form/path params, file uploads |
| Response | Status codes, body types + struct fields, `json`/`validate` tags |
| OpenAPI | Paths, parameters, requestBody, responses, components/schemas |

## How it works

huma uses `go/ast` and `go/types` to statically analyze Gin source code. No runtime, no reflection, no instrumentation.

1. **Route extraction** — Detects `r.GET("/path", handler)`, `r.Group("/prefix")`, `r.Use(mw)` patterns. Tracks group prefix accumulation and middleware inheritance. Handles `gin.IRouter` parameters and string concatenation paths (`options.BaseURL + "/path"`).

2. **Handler analysis** — Finds the actual `*ast.FuncDecl` for each handler via `go/types`. Scans the function body for `c.ShouldBindJSON`, `c.Query`, `c.JSON`, etc. Follows 1-depth calls that pass `*gin.Context` to wrapper functions (e.g., `middleware.Success(c, data)`).

3. **Type resolution** — Resolves binding/response variables to their Go types via `go/types`. Extracts struct fields with `json` and `validate` tags. Handles nested structs, pointers, slices, maps, and embedded fields. Marks `gin.H` responses as `partial` confidence.

4. **Caller-side argument propagation** — When a handler calls a wrapper function, resolves the caller's arguments (not just the callee's internals). This recovers actual status codes from `http.StatusBadRequest` and actual response data types that would otherwise be hidden behind wrapper types like `Envelope`.

5. **OpenAPI generation** — Converts extracted data to OpenAPI 3.0 YAML. Named types become `$ref` entries in `components/schemas`. Gin `:param` paths become `{param}`. `gin.H` responses get `x-schema-confidence: partial`.

## Flags

Note: flags must come **before** the project path.

```
huma scan [flags] [project-root]

  --openapi   Output OpenAPI 3.0 YAML
  --json      Output JSON
  -o string   Write to file instead of stdout
```

## Limitations

- **Dynamic routes**: `r.GET(variable, handler)` where the path is not a string literal cannot be extracted.
- **Interface dispatch**: oapi-codegen `RegisterHandlers` wraps handlers behind an interface; huma skips `.gen.go` files since those already have OpenAPI specs.
- **`gin.H` responses**: `map[string]any` — keys are extracted, value types are best-effort. Marked with `x-schema-confidence: partial`.
- **External package wrappers**: 1-depth call tracking only works for functions within the project (`./...` scope). Third-party wrapper functions cannot be followed.
- **DDL/sqlc extraction**: Planned, not yet implemented.

## License

MIT
