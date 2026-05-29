//ff:func feature=prisma type=command control=sequence
//ff:what prisma 서브커맨드 실행 (schema.prisma -> canonical CREATE TABLE)
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/park-jun-woo/codistill/internal/ddl"
	"github.com/park-jun-woo/codistill/internal/prisma"
)

func runPrisma(args []string) {
	fs := flag.NewFlagSet("prisma", flag.ExitOnError)
	outDir := fs.String("o", "", "output directory (one .sql file per table)")
	fs.Parse(args)

	path := "."
	if fs.NArg() > 0 {
		path = fs.Arg(0)
	}

	tables, err := prisma.Parse(path)
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
