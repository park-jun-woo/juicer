//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildViewSetEndpoints_NoMethodsNoActions 테스트
package django

import "testing"

func TestBuildViewSetEndpoints_NoMethodsNoActions(t *testing.T) {
	vs := &viewsetInfo{name: "Bare", parents: nil, file: "v.py"}
	eps := buildViewSetEndpoints(routerRegistration{prefix: "x"}, vs, nil)
	if len(eps) != 0 {
		t.Fatalf("expected no endpoints, got %d", len(eps))
	}
}
