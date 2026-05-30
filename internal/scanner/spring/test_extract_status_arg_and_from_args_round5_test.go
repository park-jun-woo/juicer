//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractStatusArgAndFromArgs_Round5 테스트
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestExtractStatusArgAndFromArgs_Round5(t *testing.T) {
	root, src := sParse(t, `class C { void m() { ResponseEntity.status(HttpStatus.CREATED).build(); } }`)
	var inv *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if inv != nil {
			return
		}
		if n.Type() == "method_invocation" {
			name := n.ChildByFieldName("name")
			if name != nil && nodeText(name, src) == "status" {
				inv = n
			}
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if inv == nil {
		t.Fatal("no status invocation")
	}
	got := extractStatusArg(inv, src)
	if got == "" {
		t.Fatalf("status arg empty")
	}

	root2, src2 := sParse(t, `@ResponseStatus(HttpStatus.CREATED) class C {}`)
	ann := sFirst(t, root2, "annotation")
	args := annotationArgs(ann, src2)
	if args == nil {
		t.Fatal("no annotation args")
	}
	if s := extractStatusFromArgs(args, src2); s == "" {
		t.Fatalf("extractStatusFromArgs empty")
	}
}
