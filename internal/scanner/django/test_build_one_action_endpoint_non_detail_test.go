//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildOneActionEndpoint_NonDetail 테스트
package django

import "testing"

func TestBuildOneActionEndpoint_NonDetail(t *testing.T) {
	vs := &viewsetInfo{name: "UserViewSet", file: "views.py"}
	ai := actionInfo{name: "recent", line: 9, detail: false}
	ep := buildOneActionEndpoint("users/", "recent", "GET", ai, vs)

	if ep.Path != "/users/recent" {
		t.Errorf("path = %q, want /users/recent", ep.Path)
	}
	if ep.Request != nil {
		t.Fatal("expected no path params for non-detail action")
	}
}
