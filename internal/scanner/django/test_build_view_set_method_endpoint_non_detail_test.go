//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildViewSetMethodEndpoint_NonDetail 테스트
package django

import "testing"

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
