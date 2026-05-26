//ff:func feature=scan type=test control=sequence
//ff:what TestBuildFuncIndex_Nil 테스트
package scanner

import "testing"

func TestBuildFuncIndex_Nil(t *testing.T) {
	idx := buildFuncIndex(nil)
	if idx == nil {
		t.Fatal("expected non-nil")
	}
}
