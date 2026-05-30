//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestIsJSONFieldExpr_NotFieldExpr 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestIsJSONFieldExpr_NotFieldExpr(t *testing.T) {

	if isJSONFieldExpr(nil, nil) {
		t.Fatal("expected false for nil node")
	}
	src := []byte(`fn f() { y; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	var id *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if id != nil {
			return
		}
		if n.Type() == "identifier" && nodeText(n, src) == "y" {
			id = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if isJSONFieldExpr(id, src) {
		t.Fatal("expected false for identifier node")
	}
}
