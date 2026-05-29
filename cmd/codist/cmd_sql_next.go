//ff:func feature=ratchet type=command control=sequence
//ff:what sql next 서브커맨드 빌더 — 플래그 등록 후 RunE에서 runSQLNext에 위임한다
package main

import (
	"github.com/spf13/cobra"
)

func newSQLNextCmd() *cobra.Command {
	var repoDir string
	var queriesDir string
	cmd := &cobra.Command{
		Use:   "next",
		Short: "Advance the ratchet session to the next method",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSQLNext(repoDir, queriesDir)
		},
	}
	cmd.Flags().StringVar(&repoDir, "repo", "", "repository directory (required on first run)")
	cmd.Flags().StringVar(&queriesDir, "queries", "", "sqlc queries directory (required on first run)")
	return cmd
}
