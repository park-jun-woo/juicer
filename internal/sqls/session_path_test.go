package sqls

import "testing"

func TestSessionPath_ReturnsPath(t *testing.T) {
	got := sessionPath()
	if got == "" {
		t.Fatal("expected non-empty")
	}
}
