//ff:func feature=scan type=test control=sequence topic=actix
//ff:what findParamType — 함수 파라미터의 타입 노드 탐색을 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstParam(root *sitter.Node) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "parameter" {
			found = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return found
}

func TestFindParamType_Generic(t *testing.T) {
	src := []byte(`fn f(body: web::Json<User>) {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	p := firstParam(root)
	if p == nil {
		t.Fatal("no parameter")
	}
	ty := findParamType(p)
	if ty == nil {
		t.Fatal("expected a type node")
	}
	if ty.Type() != "generic_type" {
		t.Errorf("type kind = %s, want generic_type", ty.Type())
	}
}

func TestFindParamType_NotFound(t *testing.T) {
	// An empty block has no matching type child -> nil.
	src := []byte(`fn f() {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	var block *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if block != nil {
			return
		}
		if n.Type() == "block" {
			block = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if block == nil {
		t.Fatal("no block")
	}
	if ty := findParamType(block); ty != nil {
		t.Fatalf("expected nil, got %s", ty.Type())
	}
}
