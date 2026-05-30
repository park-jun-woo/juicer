//ff:func feature=scan type=test control=sequence
//ff:what TestBuildFuncIndex_Empty 테스트
package fiber

import "testing"

func TestBuildFuncIndex_Empty(t *testing.T) {
	idx := buildFuncIndex(nil)
	if idx == nil || idx.byPos == nil || idx.byName == nil || idx.astStructs == nil {
		t.Fatal("expected initialized index maps")
	}
	if len(idx.byPos) != 0 {
		t.Errorf("expected empty byPos, got %d", len(idx.byPos))
	}
}
