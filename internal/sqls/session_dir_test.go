package sqls

import "testing"

func TestSessionDir_ReturnsDir(t *testing.T) {
	got := SessionDir()
	if got != ".huma" {
		t.Fatalf("expected .huma, got %s", got)
	}
}
