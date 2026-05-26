//ff:func feature=scan type=test control=sequence
//ff:what TestFindMappingIndex_Nil 테스트
package scanner

import "testing"

func TestFindMappingIndex_Nil(t *testing.T) {
	if idx := findMappingIndex(nil, "key"); idx != -1 {
		t.Fatalf("expected -1, got %d", idx)
	}
}
