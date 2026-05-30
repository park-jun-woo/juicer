//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildViewSetActionEndpoints_DefaultName 테스트
package django

import "testing"

func TestBuildViewSetActionEndpoints_DefaultName(t *testing.T) {
	vs := &viewsetInfo{name: "UserViewSet", file: "v.py"}

	ai := actionInfo{name: "recent", methods: []string{"GET", "POST"}, detail: false}
	eps := buildViewSetActionEndpoints("users/", ai, vs)
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
	if eps[0].Path != "/users/recent" {
		t.Errorf("path = %q, want /users/recent", eps[0].Path)
	}
}
