//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what assignOneArg 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstNodeOfType(t *testing.T, src, typ string) (*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(src))
	ns := findAllByType(fi.Root, typ)
	if len(ns) == 0 {
		t.Fatalf("no %s node in %q", typ, src)
	}
	return ns[0], fi.Src
}

func TestAssignOneArg_Object(t *testing.T) {
	n, src := firstNodeOfType(t, "const x = { a: 1 };\n", "object")
	ri := &routeInfo{}
	assignOneArg(ri, n, src)
	if ri.Schema != n {
		t.Fatal("object arg should set Schema")
	}
}

func TestAssignOneArg_Identifier(t *testing.T) {
	n, src := firstNodeOfType(t, "f(handler);\n", "identifier")
	ri := &routeInfo{}
	assignOneArg(ri, n, src)
	if ri.Handler == "" {
		t.Fatal("identifier arg should set Handler")
	}
}

func TestAssignOneArg_Arrow(t *testing.T) {
	n, src := firstNodeOfType(t, "const x = () => 1;\n", "arrow_function")
	ri := &routeInfo{}
	assignOneArg(ri, n, src)
	if ri.Handler != "(anonymous)" {
		t.Fatalf("arrow should set anonymous handler, got %q", ri.Handler)
	}
}

func TestAssignOneArg_Default(t *testing.T) {
	// a string node matches none of the cases -> no change
	n, src := firstNodeOfType(t, `const x = "lit";`+"\n", "string")
	ri := &routeInfo{Handler: "orig"}
	assignOneArg(ri, n, src)
	if ri.Handler != "orig" || ri.Schema != nil {
		t.Fatalf("default case should leave ri unchanged, got %v", ri)
	}
}
