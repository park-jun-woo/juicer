//ff:func feature=scan type=test control=sequence
//ff:what TestStatusDescription_200 테스트
package scanner

import "testing"

func TestStatusDescription_200(t *testing.T) {
	if statusDescription("200") != "OK" {
		t.Fatal("expected OK")
	}
}

