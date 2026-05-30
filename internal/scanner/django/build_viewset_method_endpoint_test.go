//ff:func feature=scan type=test control=sequence topic=django
//ff:what buildViewSetMethodEndpoint — detail/비detail CRUD 메서드 엔드포인트를 검증
package django

import "testing"

func TestBuildViewSetMethodEndpoint_Detail(t *testing.T) {
	vs := &viewsetInfo{name: "UserViewSet", file: "v.py"}
	am := actionMethod{action: "retrieve", method: "GET", detail: true}
	ep := buildViewSetMethodEndpoint("/users", am, vs, map[string]serializerInfo{})

	if ep.Path != "/users/{pk}" {
		t.Errorf("path = %q, want /users/{pk}", ep.Path)
	}
	if ep.Request == nil || len(ep.Request.PathParams) != 1 {
		t.Fatalf("expected pk path param, got %+v", ep.Request)
	}
	if ep.Handler != "UserViewSet.retrieve" {
		t.Errorf("handler = %q", ep.Handler)
	}
}

func TestBuildViewSetMethodEndpoint_NonDetail(t *testing.T) {
	vs := &viewsetInfo{name: "UserViewSet", file: "v.py"}
	am := actionMethod{action: "list", method: "GET", detail: false}
	ep := buildViewSetMethodEndpoint("/users", am, vs, map[string]serializerInfo{})

	if ep.Path != "/users" {
		t.Errorf("path = %q, want /users", ep.Path)
	}
	if ep.Request != nil {
		t.Fatal("expected no path params for list")
	}
}
