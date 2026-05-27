//ff:func feature=sql type=command control=sequence
//ff:what sql 서브커맨드 실행 — ratchet 모드 또는 기존 스켈레톤 출력
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/park-jun-woo/codistill/internal/sqls"
)

func runSQL(args []string) {
	if len(args) > 0 {
		if handled := handleSQLSubcommand(args); handled {
			return
		}
	}

	// Existing Phase 008 behavior: skeleton output
	fs := flag.NewFlagSet("sql", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "output JSON (default YAML)")
	outFile := fs.String("o", "", "output file path")
	fs.Parse(args)

	dir := "."
	if fs.NArg() > 0 {
		dir = fs.Arg(0)
	}

	result, err := sqls.Extract(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	var output []byte
	if *jsonOut {
		output, err = sqls.RenderJSON(result)
	} else {
		output, err = sqls.RenderYAML(result)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if *outFile != "" {
		if err := os.WriteFile(*outFile, output, 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "error writing file: %v\n", err)
			os.Exit(1)
		}
	} else {
		os.Stdout.Write(output)
	}
}
