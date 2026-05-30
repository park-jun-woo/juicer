//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what schemaPairs 테스트 헬퍼
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

// schemaPairs returns the pair nodes of the first object literal.
func schemaPairs(t *testing.T, objSrc string) ([]*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte("const x = "+objSrc+";\n"))
	objs := findAllByType(fi.Root, "object")
	if len(objs) == 0 {
		t.Fatal("no object")
	}
	return childrenOfType(objs[0], "pair"), fi.Src
}
