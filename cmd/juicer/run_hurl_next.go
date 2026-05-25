//ff:func feature=hurl type=command control=sequence
//ff:what hurl next 서브커맨드의 플래그 파싱 및 실행
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/park-jun-woo/juicer/hurls"
)

func runHurlNext(args []string) {
	fs := flag.NewFlagSet("hurl next", flag.ExitOnError)
	host := fs.String("host", "", "target host URL")
	testsDir := fs.String("tests", "", "hurl tests directory")
	repoDir := fs.String("repo", "", "Go source repository directory")
	fs.Parse(args)

	if err := hurls.RunNext(*host, *testsDir, *repoDir); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
