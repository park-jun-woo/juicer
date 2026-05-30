//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildFuncViewEndpoints_NoMethods 테스트
package django

import "testing"

func TestBuildFuncViewEndpoints_NoMethods(t *testing.T) {
	entry := urlEntry{pattern: "health/"}
	fv := &funcViewInfo{name: "health", methods: nil}
	eps := buildFuncViewEndpoints(entry, fv)
	if len(eps) != 0 {
		t.Fatalf("expected 0 endpoints for no methods, got %d", len(eps))
	}
}
