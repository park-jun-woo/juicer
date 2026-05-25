package scanner

import "testing"

func TestBuildFuncIndex_Nil(t *testing.T) {
	idx := buildFuncIndex(nil)
	if idx == nil {
		t.Fatal("expected non-nil")
	}
}
