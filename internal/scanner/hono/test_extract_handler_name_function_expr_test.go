//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractHandlerName_FunctionExpr 테스트
package hono

import "testing"

func TestExtractHandlerName_FunctionExpr(t *testing.T) {

	n, fi := lastArgOf(t, `app.get("/x", function () {});`)
	if got := extractHandlerName(n, fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
