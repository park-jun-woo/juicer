//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractHandlerName_Identifier 테스트
package hono

import "testing"

func TestExtractHandlerName_Identifier(t *testing.T) {
	n, fi := lastArgOf(t, `app.get("/x", handler);`)
	if got := extractHandlerName(n, fi.Src); got != "handler" {
		t.Fatalf("got %q", got)
	}
}
