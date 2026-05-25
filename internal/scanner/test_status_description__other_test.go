//ff:func feature=scan type=extract control=sequence
//ff:what TestStatusDescription_Other 테스트
package scanner

import "testing"

func TestStatusDescription_Other(t *testing.T) {
	if statusDescription("418") != "Response" {
		t.Fatal("expected Response")
	}
}
