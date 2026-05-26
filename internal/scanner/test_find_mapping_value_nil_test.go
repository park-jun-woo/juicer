//ff:func feature=scan type=test control=sequence
//ff:what TestFindMappingValue_Nil 테스트
package scanner

import "testing"

func TestFindMappingValue_Nil(t *testing.T) {
	result := findMappingValue(nil, "key")
	if result != nil {
		t.Fatal("expected nil")
	}
}
