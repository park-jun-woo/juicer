//ff:func feature=ratchet type=command control=sequence
//ff:what sql next 서브커맨드의 플래그 파싱 및 실행
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/park-jun-woo/juicer/internal/sqls"
)

func runSQLNext(args []string) {
	fs := flag.NewFlagSet("sql next", flag.ExitOnError)
	repoDir := fs.String("repo", "", "repository directory")
	queriesDir := fs.String("queries", "", "sqlc queries directory")
	fs.Parse(args)

	if err := sqls.RunNext(*repoDir, *queriesDir); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
