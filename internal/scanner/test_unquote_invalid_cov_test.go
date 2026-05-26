//ff:func feature=scan type=test control=sequence
//ff:what TestUnquote_InvalidCov 테스트
package scanner

import "testing"

func TestUnquote_InvalidCov(t *testing.T) {
	got := unquote(`"unclosed`)
	if got != "unclosed" {
		t.Fatalf("expected unclosed, got %s", got)
	}
}
