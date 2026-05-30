//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what applyArrayItems 테스트
package fastify

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

// firstObject parses a TS expression `const x = <obj>` and returns the object node.
func firstObject(t *testing.T, objSrc string) (*sitter.Node, []byte) {
	t.Helper()
	src := []byte("const x = " + objSrc + ";\n")
	fi := mustParse(t, src)
	objs := findAllByType(fi.Root, "object")
	if len(objs) == 0 {
		t.Fatal("no object node")
	}
	return objs[0], fi.Src
}

func TestApplyArrayItems_PrimitiveItems(t *testing.T) {
	obj, src := firstObject(t, `{ type: "array", items: { type: "string" } }`)
	f := &scanner.Field{}
	applyArrayItems(f, obj, src)
	if f.Type != "string[]" {
		t.Fatalf("expected string[], got %q", f.Type)
	}
}

func TestApplyArrayItems_ObjectItems(t *testing.T) {
	obj, src := firstObject(t, `{ type: "array", items: { type: "object", properties: { id: { type: "integer" } } } }`)
	f := &scanner.Field{}
	applyArrayItems(f, obj, src)
	if len(f.Fields) != 1 || f.Fields[0].Name != "id" {
		t.Fatalf("expected nested id field, got %v", f.Fields)
	}
}

func TestApplyArrayItems_NoItems(t *testing.T) {
	obj, src := firstObject(t, `{ type: "array" }`)
	f := &scanner.Field{Type: "orig"}
	applyArrayItems(f, obj, src)
	if f.Type != "orig" || len(f.Fields) != 0 {
		t.Fatalf("expected unchanged when items missing, got %v", f)
	}
}

func TestApplyArrayItems_ItemsNotObject(t *testing.T) {
	// items is a string, not an object literal -> early return
	obj, src := firstObject(t, `{ type: "array", items: "nope" }`)
	f := &scanner.Field{Type: "orig"}
	applyArrayItems(f, obj, src)
	if f.Type != "orig" {
		t.Fatalf("expected unchanged when items not object, got %q", f.Type)
	}
}
