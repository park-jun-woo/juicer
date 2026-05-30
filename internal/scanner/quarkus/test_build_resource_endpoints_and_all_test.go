//ff:func feature=scan type=test control=iteration dimension=1 topic=quarkus
//ff:what TestBuildResourceEndpointsAndAll 테스트
package quarkus

import "testing"

func TestBuildResourceEndpointsAndAll(t *testing.T) {
	fi := qFileInfo(t, sampleResource)
	resources := extractResources(fi)
	eps, _ := buildAllEndpoints(resources, "/abs")
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
