//ff:func feature=hurl type=render control=sequence
//ff:what TestFormatSlice_Empty 테스트
package hurls

import "testing"

func TestFormatSlice_Empty(t *testing.T) {
	got := formatSlice(nil)
	if got != "[]" {
		t.Fatalf("got %q", got)
	}
}
