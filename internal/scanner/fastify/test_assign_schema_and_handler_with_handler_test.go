//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestAssignSchemaAndHandler_WithHandler 테스트
package fastify

import "testing"

func TestAssignSchemaAndHandler_WithHandler(t *testing.T) {

	nodes, src := argChildren(t, `app.get("/p", { schema: {} }, myHandler);`+"\n")
	ri := &routeInfo{}
	assignSchemaAndHandler(ri, nodes, src)
	if ri.Schema == nil {
		t.Error("expected schema set")
	}
	if ri.Handler != "myHandler" {
		t.Errorf("expected myHandler, got %q", ri.Handler)
	}
}
