//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractHandlerAndMiddleware_TooFewArgs 테스트
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
