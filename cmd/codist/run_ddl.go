//ff:func feature=ddl type=command control=sequence
//ff:what ddl 서브커맨드 실행
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/park-jun-woo/codistill/internal/ddl"
)

func runDDL(args []string) {
	fs := flag.NewFlagSet("ddl", flag.ExitOnError)
	outDir := fs.String("o", "", "output directory (one .sql file per table)")
	fs.Parse(args)

	dir := "."
	if fs.NArg() > 0 {
		dir = fs.Arg(0)
	}

	tables, err := ddl.Parse(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	if *outDir != "" {
		if err := ddl.WriteFiles(tables, *outDir); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Print(ddl.Render(tables))
	}
}
