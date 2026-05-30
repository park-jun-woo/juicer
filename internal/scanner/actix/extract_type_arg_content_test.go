//ff:func feature=scan type=test control=sequence topic=actix
//ff:what extractTypeArgContent — 단일/복수 타입 인자 분기를 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstTypeArguments(root *sitter.Node) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "type_arguments" {
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

func TestExtractTypeArgContent_Single(t *testing.T) {
	src := []byte(`fn f(x: web::Json<User>) {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	ta := firstTypeArguments(root)
	if ta == nil {
		t.Fatal("no type_arguments found")
	}
	if got := extractTypeArgContent(ta, src); got != "User" {
		t.Fatalf("extractTypeArgContent = %q, want User", got)
	}
}

func TestExtractTypeArgContent_Multiple(t *testing.T) {
	// More than one named type -> falls through to joinTypeArgTokens.
	src := []byte(`fn f(x: HashMap<String, i32>) {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	ta := firstTypeArguments(root)
	if ta == nil {
		t.Fatal("no type_arguments found")
	}
	got := extractTypeArgContent(ta, src)
	// joinTypeArgTokens output should contain both type names.
	if got == "String" || got == "" {
		t.Fatalf("unexpected single-type result for multi-arg: %q", got)
	}
}
