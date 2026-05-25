//ff:func feature=scan type=extract control=sequence
//ff:what TestStatusDescription_Unknown 테스트
package scanner

import "testing"

func TestStatusDescription_Unknown(t *testing.T) {
	if statusDescription("(unknown)") != "Error" {
		t.Fatal("expected Error")
	}
}
