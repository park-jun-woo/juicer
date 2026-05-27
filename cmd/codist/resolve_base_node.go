//ff:func feature=scan type=command control=sequence
//ff:what base spec 경로를 해석하고 yaml.Node로 로드한다
package main

import (
	"fmt"
	"os"

	"github.com/park-jun-woo/codistill/internal/scanner"
	"gopkg.in/yaml.v3"
)

func resolveBaseNode(baseFile string, root string) *yaml.Node {
	basePath := baseFile
	if basePath == "" {
		basePath = scanner.FindBaseSpec(root)
	}
	if basePath == "" {
		return nil
	}
	bn, err := scanner.LoadBaseSpec(basePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: cannot load base spec %s: %v\n", basePath, err)
		return nil
	}
	return bn
}
