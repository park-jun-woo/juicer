//ff:func feature=scan type=command control=sequence
//ff:what scan 서브커맨드 실행
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/park-jun-woo/juicer/scanner"
)

func runScan(args []string) {
	fs := flag.NewFlagSet("scan", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "output JSON")
	openapiOut := fs.Bool("openapi", false, "output OpenAPI 3.0 YAML")
	outFile := fs.String("o", "", "output file path")
	fs.Parse(args)

	root := "."
	if fs.NArg() > 0 {
		root = fs.Arg(0)
	}

	result, err := scanner.Scan(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	format := scanner.FormatYAML
	if *openapiOut {
		format = scanner.FormatOpenAPI
	} else if *jsonOut {
		format = scanner.FormatJSON
	}

	output, err := scanner.Render(result, format)
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
