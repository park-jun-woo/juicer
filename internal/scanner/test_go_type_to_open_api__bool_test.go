//ff:func feature=scan type=extract control=sequence
//ff:what TestGoTypeToOpenAPI_Bool 테스트
package scanner

import "testing"

func TestGoTypeToOpenAPI_Bool(t *testing.T) {
	if goTypeToOpenAPI("bool") != "boolean" {
		t.Fatal("expected boolean")
	}
}
