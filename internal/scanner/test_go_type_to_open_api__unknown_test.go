//ff:func feature=scan type=extract control=sequence
//ff:what TestGoTypeToOpenAPI_Unknown 테스트
package scanner

import "testing"

func TestGoTypeToOpenAPI_Unknown(t *testing.T) {
	if goTypeToOpenAPI("CustomType") != "object" {
		t.Fatal("expected object")
	}
}
