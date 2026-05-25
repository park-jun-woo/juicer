package hurls

import "testing"

func TestFormatSlice(t *testing.T) {
	got := formatSlice([]string{"a", "b", "c"})
	if got != "[a, b, c]" {
		t.Fatalf("got %q", got)
	}
}

func TestFormatSlice_Empty(t *testing.T) {
	got := formatSlice(nil)
	if got != "[]" {
		t.Fatalf("got %q", got)
	}
}
