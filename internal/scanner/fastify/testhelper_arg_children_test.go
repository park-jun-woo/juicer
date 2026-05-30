//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what argChildren 테스트 헬퍼
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

// argChildren returns the non-punctuation children of the first arguments node.
func argChildren(t *testing.T, callSrc string) ([]*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(callSrc))
	args := findAllByType(fi.Root, "arguments")
	if len(args) == 0 {
		t.Fatal("no arguments node")
	}
	var nodes []*sitter.Node
	a := args[0]
	for i := 0; i < int(a.NamedChildCount()); i++ {
		nodes = append(nodes, a.NamedChild(i))
	}
	return nodes, fi.Src
}
