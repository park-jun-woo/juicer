//ff:func feature=scan type=test control=sequence topic=actix
//ff:what macroHandlerName — function_item 핸들러명 추출 분기를 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestMacroHandlerName(t *testing.T) {
	src := []byte(`async fn get_user() {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fn := findFuncByName(root, src, "get_user")
	if fn == nil {
		t.Fatal("function not found")
	}
	if got := macroHandlerName(fn, src); got != "get_user" {
		t.Fatalf("macroHandlerName = %q, want get_user", got)
	}
}

func TestMacroHandlerName_NoIdentifier(t *testing.T) {
	// A block node has no identifier child -> "".
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
	if got := macroHandlerName(block, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
