//ff:func feature=scan type=test control=sequence topic=django
//ff:what applyActionKeywords — @action 키워드 인자 적용 분기를 검증
package django

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstArgumentList(root *sitter.Node) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "argument_list" {
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

func TestApplyActionKeywords(t *testing.T) {
	src := []byte(`x = action(detail=True, methods=['get', 'post'], url_path='custom')` + "\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	if args == nil {
		t.Fatal("no argument_list")
	}
	ai := &actionInfo{}
	applyActionKeywords(ai, args, src)

	if !ai.detail {
		t.Error("expected detail true")
	}
	if len(ai.methods) != 2 {
		t.Errorf("methods = %v, want 2", ai.methods)
	}
	if ai.urlPath != "custom" {
		t.Errorf("urlPath = %q, want custom", ai.urlPath)
	}
}

func TestApplyActionKeywords_NonKeywordArgs(t *testing.T) {
	// Positional args (no keyword_argument) -> nothing applied.
	src := []byte(`x = action('a', 'b')` + "\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	ai := &actionInfo{}
	applyActionKeywords(ai, args, src)
	if ai.detail || len(ai.methods) != 0 || ai.urlPath != "" {
		t.Errorf("expected nothing applied, got %+v", ai)
	}
}
