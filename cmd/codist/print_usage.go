//ff:func feature=scan type=render control=sequence
//ff:what CLI 사용법 출력
package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `codist - Extract structured specs from web framework source code

Usage:
  codist scan [project-root] [flags]
  codist ddl [migrations-dir] [flags]
  codist sql [repository-dir] [flags]
  codist sql next [--repo DIR --queries DIR]
  codist sql status
  codist sql list
  codist sql skip
  codist sql reset
  codist version

scan flags:
  --framework string  framework to scan (gogin, nestjs, fastapi, express, spring, supafunc)
  --json              output JSON
  --openapi           output OpenAPI 3.0 YAML
  -o string           output file path

ddl flags:
  -o string   output directory (one .sql file per table)

sql flags:
  --json      output JSON (default YAML)
  -o string   output file path

sql next flags:
  --repo string      repository directory (required on first run)
  --queries string   sqlc queries directory (required on first run)

`)
}
