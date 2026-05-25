package hurls

import (
	"strings"
	"testing"
)

func TestSessionPath(t *testing.T) {
	p := sessionPath()
	if !strings.Contains(p, sessionFileName) {
		t.Fatalf("expected path to contain %s, got %s", sessionFileName, p)
	}
}
