//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestBuildURLMappedViewSetEndpoints_NoMethods 테스트
package django

import "testing"

func TestBuildURLMappedViewSetEndpoints_NoMethods(t *testing.T) {

	vs := &viewsetInfo{name: "Bare", parents: nil, file: "v.py"}
	eps := buildURLMappedViewSetEndpoints(urlEntry{pattern: "x/"}, vs, nil)
	if len(eps) != 0 {
		t.Fatalf("expected no endpoints, got %d", len(eps))
	}
}
