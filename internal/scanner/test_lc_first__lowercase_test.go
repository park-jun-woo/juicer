//ff:func feature=scan type=extract control=sequence
//ff:what TestLcFirst_Lowercase 테스트
package scanner

import "testing"

func TestLcFirst_Lowercase(t *testing.T) {
	if lcFirst("already") != "already" {
		t.Fatal("expected already")
	}
}
