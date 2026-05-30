//ff:func feature=scan type=test control=sequence topic=actix
//ff:what parseOneField — field_declaration → scanner.Field 변환 분기를 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func collectFieldDecls(root *sitter.Node) []*sitter.Node {
	var out []*sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if n.Type() == "field_declaration" {
			out = append(out, n)
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return out
}

func TestParseOneField_Plain(t *testing.T) {
	src := []byte(`struct S { name: String }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fds := collectFieldDecls(root)
	if len(fds) != 1 {
		t.Fatalf("expected 1 field decl, got %d", len(fds))
	}
	f := parseOneField(fds[0], src, nil)
	if f == nil {
		t.Fatal("expected field")
	}
	if f.Name != "name" {
		t.Errorf("name = %q", f.Name)
	}
	if f.Nullable {
		t.Error("expected not nullable")
	}
}

func TestParseOneField_OptionNullable(t *testing.T) {
	src := []byte(`struct S { bio: Option<String> }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fds := collectFieldDecls(root)
	f := parseOneField(fds[0], src, nil)
	if f == nil {
		t.Fatal("expected field")
	}
	if !f.Nullable {
		t.Error("expected nullable for Option<...>")
	}
}

func TestParseOneField_Skip(t *testing.T) {
	src := []byte(`struct S { internal: String }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fds := collectFieldDecls(root)
	// serde skip attribute -> field dropped.
	if f := parseOneField(fds[0], src, []serdeAttr{{skip: true}}); f != nil {
		t.Fatalf("expected nil for skipped field, got %+v", f)
	}
}

func TestParseOneField_NoName(t *testing.T) {
	// A block node has no field_identifier -> nil.
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
	if f := parseOneField(block, src, nil); f != nil {
		t.Fatalf("expected nil for node without field_identifier, got %+v", f)
	}
}
