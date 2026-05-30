//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what TestBuildURLMappedViewSetEndpoints 테스트
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
