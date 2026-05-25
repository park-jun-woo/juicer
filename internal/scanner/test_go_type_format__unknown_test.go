//ff:func feature=scan type=extract control=sequence
//ff:what TestGoTypeFormat_Unknown 테스트
package scanner

import "testing"

func TestGoTypeFormat_Unknown(t *testing.T) {
	got := goTypeFormat("string", Field{})
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
