//ff:func feature=scan type=extract control=sequence
//ff:what TestGoTypeToOpenAPI_Slice 테스트
package scanner

import "testing"

func TestGoTypeToOpenAPI_Slice(t *testing.T) {
	if goTypeToOpenAPI("[]int") != "array" {
		t.Fatal("expected array")
	}
}
