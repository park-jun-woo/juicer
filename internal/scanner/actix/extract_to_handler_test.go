//ff:func feature=scan type=test control=sequence topic=actix
//ff:what extractToHandler — .to(handler) 핸들러명 추출의 분기를 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

// findCallByFuncSuffix returns the first call_expression whose function text
// ends with the given suffix (e.g. ".to" / ".get").
func findCallByFuncSuffix(root *sitter.Node, src []byte, suffix string) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "call_expression" {
			fn := n.ChildByFieldName("function")
			if fn != nil {
				txt := nodeText(fn, src)
				if len(txt) >= len(suffix) && txt[len(txt)-len(suffix):] == suffix {
					found = n
					return
				}
			}
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return found
}

func TestExtractToHandler(t *testing.T) {
	src := []byte(`fn f() { web::get().to(my_handler); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".to")
	if call == nil {
		t.Fatal(".to call not found")
	}
	if got := extractToHandler(call, src); got != "my_handler" {
		t.Fatalf("extractToHandler = %q, want my_handler", got)
	}
}

func TestExtractToHandler_NoArgs(t *testing.T) {
	// web::get() has an empty arguments list -> loop finds nothing -> "".
	src := []byte(`fn f() { web::get(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::get")
	if call == nil {
		t.Fatal("get call not found")
	}
	if got := extractToHandler(call, src); got != "" {
		t.Fatalf("expected empty handler, got %q", got)
	}
}

func TestExtractToHandler_NoArgumentsNode(t *testing.T) {
	// A node with no "arguments" child (here: a bare identifier) hits the
	// args == nil early return.
	src := []byte(`fn f() { lone; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	var ident *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if ident != nil {
			return
		}
		if n.Type() == "identifier" && nodeText(n, src) == "lone" {
			ident = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if ident == nil {
		t.Fatal("identifier not found")
	}
	if got := extractToHandler(ident, src); got != "" {
		t.Fatalf("expected empty for node without arguments, got %q", got)
	}
}

func TestExtractToHandler_ClosureOnly(t *testing.T) {
	// .to(|| ...) has args but no resolvable identifier handler -> "".
	src := []byte(`fn f() { web::get().to(|| async { "x" }); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".to")
	if call == nil {
		t.Fatal(".to call not found")
	}
	if got := extractToHandler(call, src); got != "" {
		t.Fatalf("expected empty handler for closure, got %q", got)
	}
}
