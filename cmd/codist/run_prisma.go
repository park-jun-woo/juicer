//ff:func feature=prisma type=command control=sequence
//ff:what prisma 실행 (schema.prisma -> canonical CREATE TABLE)
package main

import (
	"fmt"

	"github.com/park-jun-woo/codistill/internal/ddl"
	"github.com/park-jun-woo/codistill/internal/prisma"
)

func runPrisma(path string, outDir string) error {
	tables, enums, err := prisma.Parse(path)
	if err != nil {
		return err
	}

	if outDir != "" {
		return ddl.WriteFiles(enums, tables, outDir)
	}
	fmt.Print(ddl.Render(enums, tables))
	return nil
}
