//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildRouterEndpoints_NoMatch 테스트
package django

import "testing"

func TestBuildRouterEndpoints_NoMatch(t *testing.T) {

	regs := []routerRegistration{{prefix: "x", viewsetName: "Missing"}}
	eps := buildRouterEndpoints(regs, []viewsetInfo{{name: "Other"}}, map[string]serializerInfo{})
	if len(eps) != 0 {
		t.Fatalf("expected no endpoints, got %d", len(eps))
	}
}
