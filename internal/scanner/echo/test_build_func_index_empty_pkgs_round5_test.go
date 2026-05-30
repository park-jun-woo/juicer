//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestBuildFuncIndex_EmptyPkgs_Round5 테스트
package echo

import "testing"

func TestBuildFuncIndex_EmptyPkgs_Round5(t *testing.T) {
	idx := buildFuncIndex(nil)
	if idx == nil || idx.byPos == nil || idx.byName == nil {
		t.Fatalf("index not initialized: %+v", idx)
	}
	if len(idx.byPos) != 0 {
		t.Fatalf("expected empty index, got %d", len(idx.byPos))
	}
}
