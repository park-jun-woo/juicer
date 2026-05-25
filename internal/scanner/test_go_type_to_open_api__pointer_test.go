//ff:func feature=scan type=extract control=sequence
//ff:what TestGoTypeToOpenAPI_Pointer 테스트
package scanner

import "testing"

func TestGoTypeToOpenAPI_Pointer(t *testing.T) {
	if goTypeToOpenAPI("*string") != "string" {
		t.Fatal("expected string")
	}
}
