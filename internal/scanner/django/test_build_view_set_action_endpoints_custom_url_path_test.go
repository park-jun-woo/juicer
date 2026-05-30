//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildViewSetActionEndpoints_CustomURLPath 테스트
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
