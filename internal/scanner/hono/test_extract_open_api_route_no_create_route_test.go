//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOpenAPIRoute_NoCreateRoute 테스트
package hono

import "testing"

func TestExtractOpenAPIRoute_NoCreateRoute(t *testing.T) {

	fi, call := openapiCall(t, `app.openapi(routeDef, handler);`)
	if r := extractOpenAPIRoute(call, fi.Src, "app"); r != nil {
		t.Fatalf("expected nil, got %+v", r)
	}
}
