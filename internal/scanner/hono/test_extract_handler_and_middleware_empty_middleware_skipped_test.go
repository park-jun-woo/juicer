//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractHandlerAndMiddleware_EmptyMiddlewareSkipped 테스트
package hono

import "testing"

func TestExtractHandlerAndMiddleware_EmptyMiddlewareSkipped(t *testing.T) {

	fi := mustParse(t, []byte(`app.get("/x", { a: 1 }, handler);`+"\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	nodes := collectArgNodes(args)
	h, mw := extractHandlerAndMiddleware(nodes, fi.Src)
	if h != "handler" {
		t.Fatalf("handler got %q", h)
	}
	if len(mw) != 0 {
		t.Fatalf("expected empty mw skipped, got %v", mw)
	}
}
