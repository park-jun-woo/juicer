//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestRouteToFieldExpr_NotCall 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestRouteToFieldExpr_NotCall(t *testing.T) {

	src := []byte(`fn f() { lone; }`)
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
		if n.Type() == "identifier" && nodeText(n, src) == "lone" {
			id = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if id == nil {
		t.Fatal("identifier not found")
	}
	if fe := routeToFieldExpr(id, src); fe != nil {
		t.Fatalf("expected nil for non-call node, got %v", fe)
	}
}
