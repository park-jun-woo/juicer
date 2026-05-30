//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractHandlerAndMiddleware_WithMiddleware 테스트
package hono

import "testing"

func TestExtractHandlerAndMiddleware_WithMiddleware(t *testing.T) {
	fi := mustParse(t, []byte(`app.get("/x", auth, logger, handler);`+"\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	nodes := collectArgNodes(args)
	h, mw := extractHandlerAndMiddleware(nodes, fi.Src)
	if h != "handler" {
		t.Fatalf("handler got %q", h)
	}
	if len(mw) != 2 || mw[0] != "auth" || mw[1] != "logger" {
		t.Fatalf("middleware got %v", mw)
	}
}
