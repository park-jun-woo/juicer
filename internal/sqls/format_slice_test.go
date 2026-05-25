package sqls

import "testing"

func TestFormatSlice_Basic(t *testing.T) {
	got := formatSlice([]string{"a", "b"})
	if got != "[a, b]" {
		t.Fatalf("expected [a, b], got %s", got)
	}
}

func TestFormatSlice_Empty(t *testing.T) {
	got := formatSlice(nil)
	if got != "[]" {
		t.Fatalf("expected [], got %s", got)
	}
}
