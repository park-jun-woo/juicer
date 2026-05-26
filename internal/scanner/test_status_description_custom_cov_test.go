//ff:func feature=scan type=test control=sequence
//ff:what TestStatusDescription_CustomCov 테스트
package scanner

import "testing"

func TestStatusDescription_CustomCov(t *testing.T) {
	if statusDescription("999") != "Response" {
		t.Fatal("expected Response")
	}
}
