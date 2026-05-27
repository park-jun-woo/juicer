//ff:func feature=scan type=command control=selection
//ff:what CLI 진입점 — 서브커맨드 디스패치
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "scan":
		runScan(os.Args[2:])
	case "ddl":
		runDDL(os.Args[2:])
	case "sql":
		runSQL(os.Args[2:])
	case "version":
		fmt.Printf("codist %s\n", Version)
	default:
		printUsage()
		os.Exit(1)
	}
}
