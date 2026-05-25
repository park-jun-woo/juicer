package sqls

import "testing"

func TestLoadSession_NoFile(t *testing.T) {
	_, err := LoadSession()
	if err == nil {
		t.Fatal("expected error when no session file")
	}
}
