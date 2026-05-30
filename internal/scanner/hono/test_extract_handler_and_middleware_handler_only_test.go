//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractHandlerAndMiddleware_HandlerOnly 테스트
package hono

import "testing"

func TestExtractHandlerAndMiddleware_HandlerOnly(t *testing.T) {
	fi := mustParse(t, []byte(`app.get("/x", handler);`+"\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	nodes := collectArgNodes(args)
	h, mw := extractHandlerAndMiddleware(nodes, fi.Src)
	if h != "handler" {
		t.Fatalf("handler got %q", h)
	}
	if len(mw) != 0 {
		t.Fatalf("expected no middleware, got %v", mw)
	}
}
