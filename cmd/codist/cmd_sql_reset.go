//ff:func feature=sql type=command control=sequence
//ff:what sql reset 서브커맨드 빌더 — RunE에서 sqls.RunReset에 위임한다
package main

import (
	"github.com/park-jun-woo/codistill/internal/sqls"
	"github.com/spf13/cobra"
)

func newSQLResetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "Reset the ratchet session",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return sqls.RunReset()
		},
	}
}
