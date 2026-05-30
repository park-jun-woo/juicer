//ff:func feature=scan type=test control=sequence topic=django
//ff:what callFuncName — call 노드 함수명 추출 분기를 검증
package django

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstCall(root *sitter.Node) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "call" {
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

func TestCallFuncName_Identifier(t *testing.T) {
	src := []byte("x = path('a/', view)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	call := firstCall(root)
	if call == nil {
		t.Fatal("no call")
	}
	if got := callFuncName(call, src); got != "path" {
		t.Fatalf("got %q, want path", got)
	}
}

func TestCallFuncName_Attribute(t *testing.T) {
	src := []byte("x = views.health()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	call := firstCall(root)
	if call == nil {
		t.Fatal("no call")
	}
	if got := callFuncName(call, src); got != "views.health" {
		t.Fatalf("got %q, want views.health", got)
	}
}

func TestCallFuncName_Neither(t *testing.T) {
	// A non-call node (block) has neither identifier nor attribute child here.
	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	// pass the module root which has no identifier/attribute direct child
	if got := callFuncName(root, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
