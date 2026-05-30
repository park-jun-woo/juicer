//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractMiddlewareName_Member 테스트
package hono

import "testing"

func TestExtractMiddlewareName_Member(t *testing.T) {
	got, _ := midArgOf(t, `app.get("/x", mw.auth, h);`)
	if got != "mw.auth" {
		t.Fatalf("got %q", got)
	}
}
