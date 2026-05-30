//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestAssignSchemaAndHandler_AnonymousFallback 테스트
package fastify

import "testing"

func TestAssignSchemaAndHandler_AnonymousFallback(t *testing.T) {

	nodes, src := argChildren(t, `app.get("/p");`+"\n")
	ri := &routeInfo{}
	assignSchemaAndHandler(ri, nodes, src)
	if ri.Handler != "(anonymous)" {
		t.Errorf("expected anonymous fallback, got %q", ri.Handler)
	}
}
