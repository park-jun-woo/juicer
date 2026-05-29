//ff:func feature=sql type=command control=sequence
//ff:what sql status 서브커맨드 빌더 — RunE에서 sqls.RunStatus에 위임한다
package main

import (
	"github.com/park-jun-woo/codistill/internal/sqls"
	"github.com/spf13/cobra"
)

func newSQLStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show the current ratchet session status",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return sqls.RunStatus()
		},
	}
}
