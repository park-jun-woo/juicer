//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what buildURLMappedViewSetEndpoints — as_view dict에서 메서드별 endpoint 생성을 검증
package django

import "testing"

func TestBuildURLMappedViewSetEndpointsDict(t *testing.T) {
	entry := urlEntry{
		pattern:       "items/",
		methodActions: map[string]string{"get": "list", "post": "create"},
	}
	vs := &viewsetInfo{name: "ItemViewSet", parents: []string{"BaseViewSet"}, file: "v.py"}
	eps := buildURLMappedViewSetEndpoints(entry, vs, map[string]serializerInfo{})

	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints (GET, POST), got %d", len(eps))
	}
	methods := map[string]string{}
	for _, ep := range eps {
		methods[ep.Method] = ep.Handler
	}
	if methods["GET"] != "ItemViewSet.list" {
		t.Errorf("GET handler = %q, want ItemViewSet.list", methods["GET"])
	}
	if methods["POST"] != "ItemViewSet.create" {
		t.Errorf("POST handler = %q, want ItemViewSet.create", methods["POST"])
	}
}
