//ff:func feature=scan type=test control=sequence topic=hono
//ff:what valueOfDecl 테스트 헬퍼
package hono

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func valueOfDecl(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	decls := findAllByType(fi.Root, "variable_declarator")
	if len(decls) == 0 {
		t.Fatal("no declarator")
	}
	return decls[0].ChildByFieldName("value"), fi.Src
}
