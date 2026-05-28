//ff:func feature=scan type=command control=selection
//ff:what scan 서브커맨드 실행 — 프레임워크 감지 후 해당 스캐너로 엔드포인트를 추출한다
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/park-jun-woo/codistill/internal/scanner"
	"github.com/park-jun-woo/codistill/internal/scanner/express"
	"github.com/park-jun-woo/codistill/internal/scanner/fastapi"
	"github.com/park-jun-woo/codistill/internal/scanner/gogin"
	"github.com/park-jun-woo/codistill/internal/scanner/nestjs"
	"github.com/park-jun-woo/codistill/internal/scanner/spring"
	"github.com/park-jun-woo/codistill/internal/scanner/supafunc"
)

func runScan(args []string) {
	fs := flag.NewFlagSet("scan", flag.ExitOnError)
	jsonOut := fs.Bool("json", false, "output JSON")
	openapiOut := fs.Bool("openapi", false, "output OpenAPI 3.0 YAML")
	baseFile := fs.String("base", "", "base OpenAPI spec to merge with")
	outFile := fs.String("o", "", "output file path")
	framework := fs.String("framework", "", "framework to scan (gogin, nestjs, fastapi, express, spring, supafunc)")
	fs.Parse(args)

	root := "."
	if fs.NArg() > 0 {
		root = fs.Arg(0)
	}

	fw := *framework
	if fw == "" {
		fw = scanner.DetectFramework(root)
		if fw == "" {
			fmt.Fprintf(os.Stderr, "error: could not detect framework; specify --framework\n")
			os.Exit(1)
		}
	}

	var result *scanner.ScanResult
	var err error

	switch fw {
	case "gogin":
		result, err = gogin.Scan(root)
	case "nestjs":
		result, err = nestjs.Scan(root)
	case "fastapi":
		result, err = fastapi.Scan(root)
	case "express":
		result, err = express.Scan(root)
	case "spring":
		result, err = spring.Scan(root)
	case "supafunc":
		result, err = supafunc.Scan(root)
	default:
		err = fmt.Errorf("unknown framework: %s", fw)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	var output []byte

	if *openapiOut {
		baseNode := resolveBaseNode(*baseFile, root)
		output, err = scanner.ToOpenAPI(result, baseNode)
	} else {
		format := scanner.FormatYAML
		if *jsonOut {
			format = scanner.FormatJSON
		}
		output, err = scanner.Render(result, format)
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
