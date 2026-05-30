//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildFuncViewEndpoints 테스트
package django

import "testing"

func TestBuildFuncViewEndpoints(t *testing.T) {
	entry := urlEntry{pattern: "api/users/<int:pk>/"}
	fv := &funcViewInfo{
		name:    "user_detail",
		methods: []string{"GET", "PUT"},
		file:    "views.py",
		line:    10,
	}
	eps := buildFuncViewEndpoints(entry, fv)
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
	if eps[0].Handler != "user_detail" {
		t.Errorf("handler = %q, want user_detail", eps[0].Handler)
	}
	if eps[0].Path != "/api/users/{pk}/" {
		t.Errorf("path = %q, want /api/users/{pk}/", eps[0].Path)
	}

	if eps[0].Request == nil || len(eps[0].Request.PathParams) != 1 {
		t.Fatalf("expected 1 path param, got %+v", eps[0].Request)
	}
}
