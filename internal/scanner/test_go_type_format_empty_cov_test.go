//ff:func feature=scan type=test control=sequence
//ff:what TestGoTypeFormat_EmptyCov 테스트
package scanner

import "testing"

func TestGoTypeFormat_EmptyCov(t *testing.T) {
	if got := goTypeFormat("string", Field{}); got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
