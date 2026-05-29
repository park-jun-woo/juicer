//ff:func feature=prisma type=command control=sequence
//ff:what prisma 서브커맨드 빌더 — 플래그 등록 후 RunE에서 runPrisma에 위임한다
package main

import (
	"github.com/spf13/cobra"
)

func newPrismaCmd() *cobra.Command {
	var outDir string
	cmd := &cobra.Command{
		Use:   "prisma [schema.prisma|prisma-dir]",
		Short: "Parse a Prisma schema into canonical DDL",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := "."
			if len(args) > 0 {
				path = args[0]
			}
			return runPrisma(path, outDir)
		},
	}
	cmd.Flags().StringVarP(&outDir, "output", "o", "", "output directory (one .sql file per table)")
	return cmd
}
