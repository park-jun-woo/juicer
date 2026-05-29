//ff:func feature=ddl type=command control=sequence
//ff:what ddl 실행 — migration 디렉터리를 파싱해 canonical DDL을 렌더링한다
package main

import (
	"fmt"

	"github.com/park-jun-woo/codistill/internal/ddl"
)

func runDDL(dir string, outDir string) error {
	tables, err := ddl.Parse(dir)
	if err != nil {
		return err
	}

	if outDir != "" {
		return ddl.WriteFiles(nil, tables, outDir)
	}
	fmt.Print(ddl.Render(nil, tables))
	return nil
}
