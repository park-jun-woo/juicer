//ff:func feature=scan type=test control=sequence
//ff:what TestRichness_Empty 테스트
package scanner

import "testing"

func TestRichness_Empty(t *testing.T) {
	ep := Endpoint{}
	if richness(ep) != 0 {
		t.Fatal("expected 0")
	}
}

