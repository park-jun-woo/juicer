//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractHandlerName_Call 테스트
package hono

import "testing"

func TestExtractHandlerName_Call(t *testing.T) {
	n, fi := lastArgOf(t, `app.get("/x", makeHandler());`)
	if got := extractHandlerName(n, fi.Src); got != "makeHandler" {
		t.Fatalf("got %q", got)
	}
}
