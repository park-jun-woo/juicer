//ff:func feature=scan type=test control=selection topic=django
//ff:what applyOneActionKeyword — 단일 키워드 적용 분기(unknown/url_path 비문자열)를 검증
package django

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func keywordArgs(root *sitter.Node) []*sitter.Node {
	var out []*sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if n.Type() == "keyword_argument" {
			out = append(out, n)
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return out
}

func TestApplyOneActionKeyword_UnknownKey(t *testing.T) {
	src := []byte(`x = action(permission_classes=[A])` + "\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if len(kw) == 0 {
		t.Fatal("no keyword_argument")
	}
	ai := &actionInfo{}
	applyOneActionKeyword(ai, "permission_classes", kw[0], src)
	if ai.detail || len(ai.methods) != 0 || ai.urlPath != "" {
		t.Errorf("unknown key should not modify actionInfo, got %+v", ai)
	}
}

func TestApplyOneActionKeyword_UrlPathNonString(t *testing.T) {
	// url_path bound to a non-string expression -> valNode nil -> urlPath unset.
	src := []byte(`x = action(url_path=some_var)` + "\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if len(kw) == 0 {
		t.Fatal("no keyword_argument")
	}
	ai := &actionInfo{}
	applyOneActionKeyword(ai, "url_path", kw[0], src)
	if ai.urlPath != "" {
		t.Errorf("expected empty urlPath for non-string value, got %q", ai.urlPath)
	}
}
