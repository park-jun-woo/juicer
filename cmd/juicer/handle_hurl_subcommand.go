//ff:func feature=hurl type=command control=selection
//ff:what hurl 하위 명령(next/status/list/skip/reset) 디스패치
package main

import (
	"fmt"
	"os"

	"github.com/park-jun-woo/juicer/hurls"
)

// handleHurlSubcommand dispatches hurl subcommands.
func handleHurlSubcommand(args []string) {
	switch args[0] {
	case "next":
		runHurlNext(args[1:])
	case "status":
		if err := hurls.RunStatus(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "list":
		if err := hurls.RunList(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "skip":
		if err := hurls.RunSkip(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	case "reset":
		if err := hurls.RunReset(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown hurl subcommand: %s\n", args[0])
		os.Exit(1)
	}
}
