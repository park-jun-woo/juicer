//ff:func feature=ratchet type=test control=sequence
//ff:what execSQLNext 테스트 헬퍼 — cobra 루트로 sql next를 실행하고 에러 시 os.Exit(1)
package main

import (
	"fmt"
	"os"
)

func execSQLNext(args []string) {
	c := newRootCmd()
	c.SetArgs(append([]string{"sql", "next"}, args...))
	if err := c.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
