//ff:func feature=scan type=test control=sequence topic=hono
//ff:what extractHandlerAndMiddleware 테스트
package hono

import "testing"

func TestExtractHandlerAndMiddleware_TooFewArgs(t *testing.T) {
	fi := mustParse(t, []byte(`app.get("/x");`+"\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	nodes := collectArgNodes(args)
	h, mw := extractHandlerAndMiddleware(nodes, fi.Src)
	if h != "" || mw != nil {
		t.Fatalf("expected empty/nil, got %q %v", h, mw)
	}
}

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

func TestExtractHandlerAndMiddleware_EmptyMiddlewareSkipped(t *testing.T) {
	// A middleware that yields "" (e.g. object literal) is skipped.
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
