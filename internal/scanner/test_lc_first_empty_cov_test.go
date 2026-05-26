//ff:func feature=scan type=test control=sequence
//ff:what TestLcFirst_EmptyCov 테스트
package scanner

import "testing"

func TestLcFirst_EmptyCov(t *testing.T) {
	if lcFirst("") != "" {
		t.Fatal("expected empty")
	}
}
