//ff:func feature=sql type=parse control=sequence
//ff:what TestFormatSlice_Empty 테스트
package sqls

import "testing"

func TestFormatSlice_Empty(t *testing.T) {
	got := formatSlice(nil)
	if got != "[]" {
		t.Fatalf("expected [], got %s", got)
	}
}
