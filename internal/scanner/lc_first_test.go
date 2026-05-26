//ff:func feature=scan type=test control=sequence
//ff:what TestLcFirst_Simple 테스트
package scanner

import "testing"

func TestLcFirst_Simple(t *testing.T) {
	if lcFirst("Building") != "building" {
		t.Fatal("expected building")
	}
}

