//ff:func feature=scan type=test control=sequence topic=actix
//ff:what walkMethodChain — 메서드 체인 하강 및 비메서드 호출 종료 분기를 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestWalkMethodChain_Descends(t *testing.T) {
	src := []byte(`fn f() { web::scope("/p").service(a).service(b); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	chain := findCallByFuncSuffix(root, src, ".service")
	if chain == nil {
		t.Fatal("no .service chain")
	}
	count := 0
	walkMethodChain(chain, src, "service", func(args *sitter.Node) { count++ })
	if count != 2 {
		t.Fatalf("expected 2 service callbacks, got %d", count)
	}
}

func TestWalkMethodChain_NoFieldExpr(t *testing.T) {
	// web::get() is a call_expression whose function is a scoped_identifier
	// (no field_expression child) -> walk returns immediately, 0 callbacks.
	src := []byte(`fn f() { web::get(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::get")
	if call == nil {
		t.Fatal("no get call")
	}
	count := 0
	walkMethodChain(call, src, "service", func(args *sitter.Node) { count++ })
	if count != 0 {
		t.Fatalf("expected 0 callbacks, got %d", count)
	}
}

func TestWalkMethodChain_NonCallNode(t *testing.T) {
	// A non-call_expression node never enters the loop.
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
	count := 0
	walkMethodChain(id, src, "service", func(args *sitter.Node) { count++ })
	if count != 0 {
		t.Fatalf("expected 0 callbacks, got %d", count)
	}
}
