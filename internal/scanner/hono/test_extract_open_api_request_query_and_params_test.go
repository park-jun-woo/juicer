//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOpenAPIRequest_QueryAndParams 테스트
package hono

import "testing"

func TestExtractOpenAPIRequest_QueryAndParams(t *testing.T) {
	fi, obj := createRouteObj(t, `{
  request: {
    query: z.object({ q: z.string() }),
    params: z.object({ id: z.string() })
  }
}`)
	v := extractOpenAPIRequest(obj, fi.Src)
	if len(v) != 2 {
		t.Fatalf("expected 2 validators, got %d: %+v", len(v), v)
	}
	if v[0].Target != "query" || v[1].Target != "param" {
		t.Fatalf("targets: %s %s", v[0].Target, v[1].Target)
	}
}
