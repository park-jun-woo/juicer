//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestExtractOneRoute_Success 테스트
package fastify

import "testing"

func TestExtractOneRoute_Success(t *testing.T) {
	fi, calls := routeCalls(t, `app.get("/users/:id", handler);`+"\n")
	inst := map[string]bool{"app": true}
	var got *routeInfo
	for _, c := range calls {
		if ri := extractOneRoute(c, fi.Src, inst); ri != nil {
			got = ri
		}
	}
	if got == nil {
		t.Fatal("expected route")
	}
	if got.Method != "GET" || got.Path != "/users/:id" || got.Handler != "handler" {
		t.Fatalf("route = %+v", got)
	}
}
