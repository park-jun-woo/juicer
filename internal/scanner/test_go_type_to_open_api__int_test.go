//ff:func feature=scan type=extract control=sequence
//ff:what TestGoTypeToOpenAPI_Int 테스트
package scanner

import "testing"

func TestGoTypeToOpenAPI_Int(t *testing.T) {
	if goTypeToOpenAPI("int") != "integer" {
		t.Fatal("expected integer")
	}
}
