//ff:func feature=scan type=test control=sequence topic=django
//ff:what buildURLMappedViewSetEndpoints — path() 매핑 ViewSet 엔드포인트 생성을 검증
package django

import "testing"

func TestBuildURLMappedViewSetEndpoints(t *testing.T) {
	entry := urlEntry{pattern: "items/<int:pk>/"}
	vs := &viewsetInfo{
		name:    "ItemViewSet",
		parents: []string{"ModelViewSet"},
		file:    "v.py",
	}
	eps := buildURLMappedViewSetEndpoints(entry, vs, map[string]serializerInfo{})
	if len(eps) == 0 {
		t.Fatal("expected endpoints")
	}
	for _, ep := range eps {
		if ep.Path != "/items/{pk}/" {
			t.Errorf("path = %q, want /items/{pk}/", ep.Path)
		}
		if ep.Request == nil || len(ep.Request.PathParams) != 1 {
			t.Errorf("expected pk path param, got %+v", ep.Request)
		}
	}
}

func TestBuildURLMappedViewSetEndpoints_NoMethods(t *testing.T) {
	// A ViewSet with no recognized parents yields no methods -> no endpoints.
	vs := &viewsetInfo{name: "Bare", parents: nil, file: "v.py"}
	eps := buildURLMappedViewSetEndpoints(urlEntry{pattern: "x/"}, vs, nil)
	if len(eps) != 0 {
		t.Fatalf("expected no endpoints, got %d", len(eps))
	}
}
