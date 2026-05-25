//ff:func feature=hurl type=render control=sequence
//ff:what TestFormatSlice 테스트
package hurls

import "testing"

func TestFormatSlice(t *testing.T) {
	got := formatSlice([]string{"a", "b", "c"})
	if got != "[a, b, c]" {
		t.Fatalf("got %q", got)
	}
}
