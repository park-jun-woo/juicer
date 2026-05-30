//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what assignSchemaAndHandler 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

// argChildren returns the non-punctuation children of the first arguments node.
func argChildren(t *testing.T, callSrc string) ([]*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(callSrc))
	args := findAllByType(fi.Root, "arguments")
	if len(args) == 0 {
		t.Fatal("no arguments node")
	}
	var nodes []*sitter.Node
	a := args[0]
	for i := 0; i < int(a.NamedChildCount()); i++ {
		nodes = append(nodes, a.NamedChild(i))
	}
	return nodes, fi.Src
}

func TestAssignSchemaAndHandler_WithHandler(t *testing.T) {
	// args: "/path", {schema}, handler
	nodes, src := argChildren(t, `app.get("/p", { schema: {} }, myHandler);`+"\n")
	ri := &routeInfo{}
	assignSchemaAndHandler(ri, nodes, src)
	if ri.Schema == nil {
		t.Error("expected schema set")
	}
	if ri.Handler != "myHandler" {
		t.Errorf("expected myHandler, got %q", ri.Handler)
	}
}

func TestAssignSchemaAndHandler_AnonymousFallback(t *testing.T) {
	// only the path arg, no handler -> fallback "(anonymous)"
	nodes, src := argChildren(t, `app.get("/p");`+"\n")
	ri := &routeInfo{}
	assignSchemaAndHandler(ri, nodes, src)
	if ri.Handler != "(anonymous)" {
		t.Errorf("expected anonymous fallback, got %q", ri.Handler)
	}
}
