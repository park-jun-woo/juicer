//ff:func feature=scan type=test control=sequence topic=django
//ff:what buildViewSetActionEndpoints — @action 커스텀 엔드포인트 생성 분기를 검증
package django

import "testing"

func TestBuildViewSetActionEndpoints_CustomURLPath(t *testing.T) {
	vs := &viewsetInfo{name: "UserViewSet", file: "v.py"}
	ai := actionInfo{name: "set_password", urlPath: "set-pw", methods: []string{"POST"}, detail: true}
	eps := buildViewSetActionEndpoints("users/", ai, vs)
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if eps[0].Path != "/users/{pk}/set-pw" {
		t.Errorf("path = %q, want /users/{pk}/set-pw", eps[0].Path)
	}
}

func TestBuildViewSetActionEndpoints_DefaultName(t *testing.T) {
	vs := &viewsetInfo{name: "UserViewSet", file: "v.py"}
	// No urlPath -> uses action name; two methods -> two endpoints.
	ai := actionInfo{name: "recent", methods: []string{"GET", "POST"}, detail: false}
	eps := buildViewSetActionEndpoints("users/", ai, vs)
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
	if eps[0].Path != "/users/recent" {
		t.Errorf("path = %q, want /users/recent", eps[0].Path)
	}
}
