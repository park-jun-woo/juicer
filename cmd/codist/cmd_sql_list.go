//ff:func feature=sql type=command control=sequence
//ff:what sql list 서브커맨드 빌더 — RunE에서 sqls.RunList에 위임한다
package main

import (
	"github.com/park-jun-woo/codistill/internal/sqls"
	"github.com/spf13/cobra"
)

func newSQLListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List ratchet session methods",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return sqls.RunList()
		},
	}
}
