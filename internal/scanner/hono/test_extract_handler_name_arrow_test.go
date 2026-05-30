//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractHandlerName_Arrow 테스트
package hono

import "testing"

func TestExtractHandlerName_Arrow(t *testing.T) {
	n, fi := lastArgOf(t, `app.get("/x", () => {});`)
	if got := extractHandlerName(n, fi.Src); got != "(anonymous)" {
		t.Fatalf("got %q", got)
	}
}
