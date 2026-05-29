//ff:func feature=scan type=command control=sequence
//ff:what version 서브커맨드 빌더 — ldflags 주입 버전을 출력한다
package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the codist version",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("codist %s\n", Version)
			return nil
		},
	}
}
