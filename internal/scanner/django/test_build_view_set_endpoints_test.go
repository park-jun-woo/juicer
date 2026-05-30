//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what TestBuildViewSetEndpoints 테스트
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
