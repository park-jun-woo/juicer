//ff:func feature=scan type=test control=sequence topic=django
//ff:what buildOneActionEndpoint — detail/비detail @action 엔드포인트 생성을 검증
package django

import "testing"

func TestBuildOneActionEndpoint_Detail(t *testing.T) {
	vs := &viewsetInfo{name: "UserViewSet", file: "views.py"}
	ai := actionInfo{name: "set_password", line: 5, detail: true}
	ep := buildOneActionEndpoint("users/", "set-password", "POST", ai, vs)

	if ep.Path != "/users/{pk}/set-password" {
		t.Errorf("path = %q, want /users/{pk}/set-password", ep.Path)
	}
	if ep.Handler != "UserViewSet.set_password" {
		t.Errorf("handler = %q", ep.Handler)
	}
	if ep.Request == nil || len(ep.Request.PathParams) != 1 {
		t.Fatalf("expected pk path param, got %+v", ep.Request)
	}
}

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
