//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOpenAPIRequest_SectionNoZodCalls 테스트
package hono

import "testing"

func TestExtractOpenAPIRequest_SectionNoZodCalls(t *testing.T) {

	fi, obj := createRouteObj(t, `{ request: { body: { content: {} } } }`)
	if v := extractOpenAPIRequest(obj, fi.Src); v != nil {
		t.Fatalf("expected nil, got %+v", v)
	}
}
