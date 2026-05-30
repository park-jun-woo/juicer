//ff:func feature=scan type=test control=sequence topic=hono
//ff:what lastArgOf 테스트 헬퍼
package hono

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func lastArgOf(t *testing.T, src string) (*sitter.Node, *fileInfo) {
	t.Helper()
	fi := mustParse(t, []byte(src+"\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	nodes := collectArgNodes(args)
	return nodes[len(nodes)-1], fi
}
