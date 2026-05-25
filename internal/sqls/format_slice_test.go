//ff:func feature=sql type=parse control=sequence
//ff:what TestFormatSlice_Basic 테스트
package sqls

import "testing"

func TestFormatSlice_Basic(t *testing.T) {
	got := formatSlice([]string{"a", "b"})
	if got != "[a, b]" {
		t.Fatalf("expected [a, b], got %s", got)
	}
}
