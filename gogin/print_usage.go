//ff:func feature=scan type=render control=sequence
//ff:what CLI 사용법 출력
package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `huma - Go Gin Endpoint Finder

Usage:
  huma scan [project-root] [flags]
  huma ddl [migrations-dir] [flags]
  huma sql [repository-dir] [flags]
  huma sql next [--repo DIR --queries DIR]
  huma sql status
  huma sql list
  huma sql skip
  huma sql reset
  huma hurl next [--host URL --tests DIR --repo DIR]
  huma hurl status
  huma hurl list
  huma hurl skip
  huma hurl reset
  huma version

scan flags:
  --json      output JSON
  --openapi   output OpenAPI 3.0 YAML
  -o string   output file path

ddl flags:
  -o string   output directory (one .sql file per table)

sql flags:
  --json      output JSON (default YAML)
  -o string   output file path

sql next flags:
  --repo string      repository directory (required on first run)
  --queries string   sqlc queries directory (required on first run)

hurl next flags:
  --host string    target host URL (required on first run)
  --tests string   hurl tests directory (required on first run)
  --repo string    Go source repository directory (required on first run)
`)
}
