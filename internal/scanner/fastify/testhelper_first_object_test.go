//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what firstObject 테스트 헬퍼
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

// firstObject parses a TS expression `const x = <obj>` and returns the object node.
func firstObject(t *testing.T, objSrc string) (*sitter.Node, []byte) {
	t.Helper()
	src := []byte("const x = " + objSrc + ";\n")
	fi := mustParse(t, src)
	objs := findAllByType(fi.Root, "object")
	if len(objs) == 0 {
		t.Fatal("no object node")
	}
	return objs[0], fi.Src
}
