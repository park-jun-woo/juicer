//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOpenAPIRoute_NoHandlerArg 테스트
package hono

import "testing"

func TestExtractOpenAPIRoute_NoHandlerArg(t *testing.T) {

	fi, call := openapiCall(t, `app.openapi(createRoute({ method: "get", path: "/x" }));`)
	r := extractOpenAPIRoute(call, fi.Src, "app")
	if r == nil || r.Method != "GET" || r.Path != "/x" || r.Handler != "" {
		t.Fatalf("got %+v", r)
	}
	if r.OwnerVar != "app" || r.Line != 1 {
		t.Fatalf("owner/line: %s %d", r.OwnerVar, r.Line)
	}
}
