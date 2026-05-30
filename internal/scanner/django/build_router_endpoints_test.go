//ff:func feature=scan type=test control=sequence topic=django
//ff:what buildRouterEndpoints — router 등록 ViewSet 엔드포인트 생성 분기를 검증
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

func TestBuildRouterEndpoints_NoMatch(t *testing.T) {
	// Registration references a ViewSet that is not in the viewsets list.
	regs := []routerRegistration{{prefix: "x", viewsetName: "Missing"}}
	eps := buildRouterEndpoints(regs, []viewsetInfo{{name: "Other"}}, map[string]serializerInfo{})
	if len(eps) != 0 {
		t.Fatalf("expected no endpoints, got %d", len(eps))
	}
}
