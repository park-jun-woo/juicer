//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestBuildAllEndpoints 테스트
package spring

import "testing"

func TestBuildAllEndpoints(t *testing.T) {
	fi := sFileInfo(t, sampleController)
	controllers := extractControllers(fi)
	eps, _ := buildAllEndpoints(controllers, "/abs")
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
	found := false
	for _, e := range eps {
		if e.Path == "/users/{id}" {
			found = true
		}
	}
	if !found {
		t.Fatalf("missing /users/{id}: %+v", eps)
	}
}
