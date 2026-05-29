//ff:func feature=sql type=command control=sequence
//ff:what sql 실행 — 저장소를 스캔해 SQL 스켈레톤을 YAML/JSON으로 출력한다
package main

import (
	"os"

	"github.com/park-jun-woo/codistill/internal/sqls"
)

func runSQL(dir string, jsonOut bool, outFile string) error {
	result, err := sqls.Extract(dir)
	if err != nil {
		return err
	}

	var output []byte
	if jsonOut {
		output, err = sqls.RenderJSON(result)
	} else {
		output, err = sqls.RenderYAML(result)
	}
	if err != nil {
		return err
	}

	if outFile != "" {
		return os.WriteFile(outFile, output, 0o644)
	}
	os.Stdout.Write(output)
	return nil
}
