//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractHandlerName_Member 테스트
package hono

import "testing"

func TestExtractHandlerName_Member(t *testing.T) {
	n, fi := lastArgOf(t, `app.get("/x", ctrl.handler);`)
	if got := extractHandlerName(n, fi.Src); got != "ctrl.handler" {
		t.Fatalf("got %q", got)
	}
}
