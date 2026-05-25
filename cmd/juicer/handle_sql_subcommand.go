//ff:func feature=sql type=command control=selection
//ff:what sql 하위 명령(next/status/list/skip/reset) 디스패치
package main

import (
	"fmt"
	"os"

	"github.com/park-jun-woo/juicer/internal/sqls"
)

// handleSQLSubcommand dispatches sql subcommands. Returns true if handled.
func handleSQLSubcommand(args []string) bool {
	switch args[0] {
	case "next":
		runSQLNext(args[1:])
		return true
	case "status":
		if err := sqls.RunStatus(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return true
	case "list":
		if err := sqls.RunList(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return true
	case "skip":
		if err := sqls.RunSkip(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return true
	case "reset":
		if err := sqls.RunReset(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return true
	}
	return false
}
