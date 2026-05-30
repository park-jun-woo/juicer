//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractHandlerName_Other 테스트
package hono

import "testing"

func TestExtractHandlerName_Other(t *testing.T) {
	n, fi := lastArgOf(t, `app.get("/x", { a: 1 });`)
	if got := extractHandlerName(n, fi.Src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
