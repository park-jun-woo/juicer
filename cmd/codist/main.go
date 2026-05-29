//ff:func feature=scan type=command control=sequence
//ff:what CLI 진입점 — 루트 cobra 커맨드를 실행하고 에러 시 비0 종료한다
package main

import (
	"fmt"
	"os"
)

func main() {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
