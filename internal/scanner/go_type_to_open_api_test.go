//ff:func feature=scan type=test control=sequence
//ff:what TestGoTypeToOpenAPI_String 테스트
package scanner

import "testing"

func TestGoTypeToOpenAPI_String(t *testing.T) {
	if goTypeToOpenAPI("string") != "string" {
		t.Fatal("expected string")
	}
}

