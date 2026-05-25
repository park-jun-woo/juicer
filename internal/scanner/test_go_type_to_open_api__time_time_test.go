//ff:func feature=scan type=extract control=sequence
//ff:what TestGoTypeToOpenAPI_TimeTime 테스트
package scanner

import "testing"

func TestGoTypeToOpenAPI_TimeTime(t *testing.T) {
	if goTypeToOpenAPI("time.Time") != "string" {
		t.Fatal("expected string")
	}
}
