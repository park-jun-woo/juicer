//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestExtractRoutes_EmptyPkgs_Round5 테스트
package echo

import "testing"

func TestExtractRoutes_EmptyPkgs_Round5(t *testing.T) {
	eps, hmap := extractRoutes(nil, "/root")
	if len(eps) != 0 || len(hmap) != 0 {
		t.Fatalf("expected empty, got %d %d", len(eps), len(hmap))
	}
}
