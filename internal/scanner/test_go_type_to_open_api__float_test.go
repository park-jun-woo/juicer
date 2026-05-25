//ff:func feature=scan type=extract control=sequence
//ff:what TestGoTypeToOpenAPI_Float 테스트
package scanner

import "testing"

func TestGoTypeToOpenAPI_Float(t *testing.T) {
	if goTypeToOpenAPI("float64") != "number" {
		t.Fatal("expected number")
	}
}
