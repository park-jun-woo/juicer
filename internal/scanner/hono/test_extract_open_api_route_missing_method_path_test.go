//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOpenAPIRoute_MissingMethodPath 테스트
package hono

import "testing"

func TestExtractOpenAPIRoute_MissingMethodPath(t *testing.T) {
	fi, call := openapiCall(t, `app.openapi(createRoute({ summary: "x" }), handler);`)
	if r := extractOpenAPIRoute(call, fi.Src, "app"); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
