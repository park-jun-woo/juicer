//ff:func feature=scan type=command control=sequence
//ff:what 루트 cobra 커맨드 생성 — 버전/usage 정책 설정 후 서브커맨드를 등록한다
package main

import (
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "codist",
		Short:         "Extract structured specs from web framework source code",
		Version:       Version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.SetVersionTemplate("codist {{.Version}}\n")
	root.AddCommand(newScanCmd())
	root.AddCommand(newDDLCmd())
	root.AddCommand(newPrismaCmd())
	root.AddCommand(newSQLCmd())
	root.AddCommand(newVersionCmd())
	return root
}
