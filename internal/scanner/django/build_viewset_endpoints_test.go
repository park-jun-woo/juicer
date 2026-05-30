//ff:func feature=scan type=test control=sequence topic=django
//ff:what buildViewSetEndpoints — CRUD + 커스텀 액션 엔드포인트 생성을 검증
package django

import "testing"

func TestBuildViewSetEndpoints(t *testing.T) {
	reg := routerRegistration{prefix: "users"}
	vs := &viewsetInfo{
		name:    "UserViewSet",
		parents: []string{"ModelViewSet"},
		file:    "v.py",
		actions: []actionInfo{
			{name: "set_password", methods: []string{"POST"}, detail: true},
		},
	}
	eps := buildViewSetEndpoints(reg, vs, map[string]serializerInfo{})
	if len(eps) == 0 {
		t.Fatal("expected endpoints")
	}
	// Should include the custom action endpoint.
	foundAction := false
	for _, ep := range eps {
		if ep.Handler == "UserViewSet.set_password" {
			foundAction = true
		}
		if ep.Path[0] != '/' {
			t.Errorf("path should start with /, got %q", ep.Path)
		}
	}
	if !foundAction {
		t.Error("expected custom @action endpoint present")
	}
}

func TestBuildViewSetEndpoints_NoMethodsNoActions(t *testing.T) {
	vs := &viewsetInfo{name: "Bare", parents: nil, file: "v.py"}
	eps := buildViewSetEndpoints(routerRegistration{prefix: "x"}, vs, nil)
	if len(eps) != 0 {
		t.Fatalf("expected no endpoints, got %d", len(eps))
	}
}
