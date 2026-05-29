//ff:func feature=ddl type=command control=sequence
//ff:what ddl 서브커맨드 빌더 — 플래그 등록 후 RunE에서 runDDL에 위임한다
package main

import (
	"github.com/spf13/cobra"
)

func newDDLCmd() *cobra.Command {
	var outDir string
	cmd := &cobra.Command{
		Use:   "ddl [migrations-dir]",
		Short: "Extract canonical DDL from migration files",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := "."
			if len(args) > 0 {
				dir = args[0]
			}
			return runDDL(dir, outDir)
		},
	}
	cmd.Flags().StringVarP(&outDir, "output", "o", "", "output directory (one .sql file per table)")
	return cmd
}
