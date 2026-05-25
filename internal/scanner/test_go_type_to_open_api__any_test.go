//ff:func feature=scan type=extract control=sequence
//ff:what TestGoTypeToOpenAPI_Any 테스트
package scanner

import "testing"

func TestGoTypeToOpenAPI_Any(t *testing.T) {
	if goTypeToOpenAPI("any") != "object" {
		t.Fatal("expected object")
	}
}
