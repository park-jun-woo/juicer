//ff:func feature=scan type=extract control=sequence
//ff:what TestLcFirst_Empty 테스트
package scanner

import "testing"

func TestLcFirst_Empty(t *testing.T) {
	if lcFirst("") != "" {
		t.Fatal("expected empty")
	}
}
