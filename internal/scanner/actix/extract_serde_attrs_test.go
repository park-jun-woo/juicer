//ff:func feature=scan type=test control=sequence topic=actix
//ff:what parseSerdeAttribute — #[serde(...)] 파싱 및 비serde nil 분기 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func collectAttrItems(root *sitter.Node) []*sitter.Node {
	var items []*sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if n.Type() == "attribute_item" {
			items = append(items, n)
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return items
}

func TestParseSerdeAttribute(t *testing.T) {
	root, err := parseRust([]byte(serdeAttrsSource))
	if err != nil {
		t.Fatal(err)
	}
	var renamed, def, skip bool
	for _, item := range collectAttrItems(root) {
		sa := parseSerdeAttribute(item, []byte(serdeAttrsSource))
		if sa == nil {
			continue // non-serde attr (e.g. #[derive(...)] / #[post(...)])
		}
		if sa.rename == "userName" {
			renamed = true
		}
		if sa.hasDefault {
			def = true
		}
		if sa.skip {
			skip = true
		}
	}
	if !renamed {
		t.Error("expected a serde rename = userName to be parsed")
	}
	if !def {
		t.Error("expected a serde default to be parsed")
	}
	if !skip {
		t.Error("expected a serde skip to be parsed")
	}
}

func TestParseSerdeAttribute_NonSerde(t *testing.T) {
	// A #[derive(...)] attribute is not serde -> serdeTokenTree returns nil.
	src := []byte("#[derive(Deserialize)]\nstruct X { a: i32 }\n")
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	items := collectAttrItems(root)
	if len(items) == 0 {
		t.Fatal("no attribute_item found")
	}
	if sa := parseSerdeAttribute(items[0], src); sa != nil {
		t.Fatalf("expected nil for non-serde attribute, got %+v", sa)
	}
}
