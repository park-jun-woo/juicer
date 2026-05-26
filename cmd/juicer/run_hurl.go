//ff:func feature=hurl type=command control=sequence
//ff:what hurl 서브커맨드 실행 — ratchet 모드
package main

import (
	"fmt"
	"os"
)

func runHurl(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: juicer hurl <next|status|list|skip|reset>\n")
		os.Exit(1)
	}
	handleHurlSubcommand(args)
}
