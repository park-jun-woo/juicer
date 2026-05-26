//ff:func feature=scan type=test control=sequence
//ff:what TestStatusDescription_UnknownCov 테스트
package scanner

import "testing"

func TestStatusDescription_UnknownCov(t *testing.T) {
	if statusDescription("(unknown)") != "Error" {
		t.Fatal("expected Error")
	}
}
