//ff:func feature=scan type=test control=sequence topic=actix
//ff:what webMethodFromCall — web::<method> 빌더를 HTTP 메서드로 변환하는 분기를 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestWebMethodFromCall_Match(t *testing.T) {
	src := []byte(`fn f() { web::post(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::post")
	if call == nil {
		t.Fatal("no call")
	}
	if got := webMethodFromCall(call, src); got != "POST" {
		t.Fatalf("got %q, want POST", got)
	}
}

func TestWebMethodFromCall_NotCall(t *testing.T) {
	// A non-call_expression node -> "".
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
	if got := webMethodFromCall(id, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestWebMethodFromCall_NoScopedIdentifier(t *testing.T) {
	// A method-call (field_expression function) has no scoped_identifier child.
	src := []byte(`fn f() { x.foo(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".foo")
	if call == nil {
		t.Fatal("no call")
	}
	if got := webMethodFromCall(call, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestWebMethodFromCall_NonWebScope(t *testing.T) {
	// scoped_identifier present but not web::<method> (len/prefix mismatch).
	src := []byte(`fn f() { other::get(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::get")
	if call == nil {
		t.Fatal("no call")
	}
	if got := webMethodFromCall(call, src); got != "" {
		t.Fatalf("expected empty for non-web scope, got %q", got)
	}
}

func TestWebMethodFromCall_UnknownWebMethod(t *testing.T) {
	// web::something that is not a known method builder -> "".
	src := []byte(`fn f() { web::unknownthing(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::unknownthing")
	if call == nil {
		t.Fatal("no call")
	}
	if got := webMethodFromCall(call, src); got != "" {
		t.Fatalf("expected empty for unknown web method, got %q", got)
	}
}
