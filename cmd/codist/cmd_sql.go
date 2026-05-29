//ff:func feature=sql type=command control=sequence
//ff:what sql 서브커맨드 빌더 — 플래그/하위커맨드 등록 후 RunE에서 runSQL에 위임한다
package main

import (
	"github.com/spf13/cobra"
)

func newSQLCmd() *cobra.Command {
	var jsonOut bool
	var outFile string
	cmd := &cobra.Command{
		Use:   "sql [repository-dir]",
		Short: "Extract SQL skeletons from repository code",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := "."
			if len(args) > 0 {
				dir = args[0]
			}
			return runSQL(dir, jsonOut, outFile)
		},
	}
	cmd.Flags().BoolVar(&jsonOut, "json", false, "output JSON (default YAML)")
	cmd.Flags().StringVarP(&outFile, "output", "o", "", "output file path")
	cmd.AddCommand(newSQLNextCmd())
	cmd.AddCommand(newSQLStatusCmd())
	cmd.AddCommand(newSQLListCmd())
	cmd.AddCommand(newSQLSkipCmd())
	cmd.AddCommand(newSQLResetCmd())
	return cmd
}
