//ff:func feature=ddl type=test control=sequence
//ff:what execDDL 테스트 헬퍼 — cobra 루트로 ddl을 실행하고 에러 시 os.Exit(1)
package main

import (
	"fmt"
	"os"
)

func execDDL(args []string) {
	c := newRootCmd()
	c.SetArgs(append([]string{"ddl"}, args...))
	if err := c.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
