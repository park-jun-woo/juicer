//ff:func feature=scan type=test control=sequence
//ff:what TestUnquote_Quoted 테스트
package scanner

import "testing"

func TestUnquote_Quoted(t *testing.T) {
	got := unquote(`"hello"`)
	if got != "hello" {
		t.Fatalf("expected hello, got %s", got)
	}
}

