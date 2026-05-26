//ff:func feature=scan type=render control=sequence
//ff:what CLI 사용법 출력
package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `juicer - Go Gin Endpoint Finder

Usage:
  juicer scan [project-root] [flags]
  juicer ddl [migrations-dir] [flags]
  juicer sql [repository-dir] [flags]
  juicer sql next [--repo DIR --queries DIR]
  juicer sql status
  juicer sql list
  juicer sql skip
  juicer sql reset
  juicer version

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

`)
}
