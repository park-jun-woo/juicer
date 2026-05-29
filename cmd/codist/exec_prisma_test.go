//ff:func feature=prisma type=test control=sequence
//ff:what execPrisma 테스트 헬퍼 — cobra 루트로 prisma를 실행하고 에러 시 os.Exit(1)
package main

import (
	"fmt"
	"os"
)

func execPrisma(args []string) {
	c := newRootCmd()
	c.SetArgs(append([]string{"prisma"}, args...))
	if err := c.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
