//ff:func feature=sql type=command control=sequence
//ff:what sql skip 서브커맨드 빌더 — RunE에서 sqls.RunSkip에 위임한다
package main

import (
	"github.com/park-jun-woo/codistill/internal/sqls"
	"github.com/spf13/cobra"
)

func newSQLSkipCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "skip",
		Short: "Skip the current ratchet method",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return sqls.RunSkip()
		},
	}
}
