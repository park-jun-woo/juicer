//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOpenAPIRoute_TwoArgsHandler 테스트
package hono

import "testing"

func TestExtractOpenAPIRoute_TwoArgsHandler(t *testing.T) {

	fi, call := openapiCall(t, `app.openapi(createRoute({ method: "put", path: "/p" }), updateH);`)
	r := extractOpenAPIRoute(call, fi.Src, "app")
	if r == nil || r.Method != "PUT" || r.Handler != "updateH" {
		t.Fatalf("got %+v", r)
	}
}
