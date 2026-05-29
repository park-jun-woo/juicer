//ff:func feature=scan type=test control=sequence
//ff:what execScan 테스트 헬퍼 — cobra 루트로 scan을 실행하고 에러 시 os.Exit(1)
package main

import (
	"fmt"
	"os"
)

func execScan(args []string) {
	c := newRootCmd()
	c.SetArgs(append([]string{"scan"}, args...))
	if err := c.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
