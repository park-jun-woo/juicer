//ff:func feature=scan type=render control=sequence
//ff:what scan 결과를 OpenAPI 또는 YAML/JSON으로 렌더링해 파일 또는 stdout에 기록한다
package main

import (
	"os"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func writeScanResult(result *scanner.ScanResult, root string, o scanOptions) error {
	var output []byte
	var err error

	if o.openapi {
		baseNode := resolveBaseNode(o.baseFile, root)
		output, err = scanner.ToOpenAPI(result, baseNode)
	} else {
		format := scanner.FormatYAML
		if o.jsonOut {
			format = scanner.FormatJSON
		}
		output, err = scanner.Render(result, format)
	}
	if err != nil {
		return err
	}

	if o.outFile != "" {
		return os.WriteFile(o.outFile, output, 0o644)
	}
	os.Stdout.Write(output)
	return nil
}
