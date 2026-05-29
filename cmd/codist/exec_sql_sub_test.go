//ff:func feature=sql type=test control=selection
//ff:what execSQLSub 테스트 헬퍼 — 알려진 sql 하위커맨드면 cobra로 실행(에러 시 os.Exit), 아니면 false
package main

import (
	"fmt"
	"os"
)

func execSQLSub(args []string) bool {
	switch args[0] {
	case "next", "status", "list", "skip", "reset":
		c := newRootCmd()
		c.SetArgs(append([]string{"sql"}, args...))
		if err := c.Execute(); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		return true
	default:
		return false
	}
}
