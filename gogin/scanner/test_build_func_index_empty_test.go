//ff:func feature=scan type=extract control=sequence
//ff:what TestBuildFuncIndex_Empty 테스트
package scanner

import (
	"testing"
)

func TestBuildFuncIndex_Empty(t *testing.T) {
	idx := buildFuncIndex(nil)
	if len(idx.byPos) != 0 {
		t.Errorf("expected 0 functions, got %d", len(idx.byPos))
	}
}
