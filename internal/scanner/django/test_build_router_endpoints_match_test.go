//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildRouterEndpoints_Match 테스트
package django

import "testing"

func TestBuildRouterEndpoints_Match(t *testing.T) {
	regs := []routerRegistration{
		{prefix: "users", viewsetName: "views.UserViewSet"},
	}
	viewsets := []viewsetInfo{
		{name: "UserViewSet", parents: []string{"ModelViewSet"}, file: "views.py"},
	}
	eps := buildRouterEndpoints(regs, viewsets, map[string]serializerInfo{})
	if len(eps) == 0 {
		t.Fatal("expected endpoints for matched ViewSet")
	}
}
