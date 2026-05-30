//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOpenAPIRequest_NoRequest 테스트
package hono

import "testing"

func TestExtractOpenAPIRequest_NoRequest(t *testing.T) {
	fi, obj := createRouteObj(t, `{ method: "get", path: "/x" }`)
	if v := extractOpenAPIRequest(obj, fi.Src); v != nil {
		t.Fatalf("expected nil, got %+v", v)
	}
}
