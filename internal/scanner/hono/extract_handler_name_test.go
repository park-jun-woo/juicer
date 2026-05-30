//ff:func feature=scan type=test control=sequence topic=hono
//ff:what extractHandlerName 테스트
package hono

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func lastArgOf(t *testing.T, src string) (*sitter.Node, *fileInfo) {
	t.Helper()
	fi := mustParse(t, []byte(src+"\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	nodes := collectArgNodes(args)
	return nodes[len(nodes)-1], fi
}

func TestExtractHandlerName_Identifier(t *testing.T) {
	n, fi := lastArgOf(t, `app.get("/x", handler);`)
	if got := extractHandlerName(n, fi.Src); got != "handler" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractHandlerName_Member(t *testing.T) {
	n, fi := lastArgOf(t, `app.get("/x", ctrl.handler);`)
	if got := extractHandlerName(n, fi.Src); got != "ctrl.handler" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractHandlerName_Call(t *testing.T) {
	n, fi := lastArgOf(t, `app.get("/x", makeHandler());`)
	if got := extractHandlerName(n, fi.Src); got != "makeHandler" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractHandlerName_Arrow(t *testing.T) {
	n, fi := lastArgOf(t, `app.get("/x", () => {});`)
	if got := extractHandlerName(n, fi.Src); got != "(anonymous)" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractHandlerName_FunctionExpr(t *testing.T) {
	// tree-sitter emits "function_expression" (not "function") for this form,
	// so it falls through to the default empty return.
	n, fi := lastArgOf(t, `app.get("/x", function () {});`)
	if got := extractHandlerName(n, fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractHandlerName_Other(t *testing.T) {
	n, fi := lastArgOf(t, `app.get("/x", { a: 1 });`)
	if got := extractHandlerName(n, fi.Src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
