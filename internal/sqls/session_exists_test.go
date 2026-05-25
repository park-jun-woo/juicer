package sqls

import "testing"

func TestSessionExists_NoFile(t *testing.T) {
	if SessionExists() {
		t.Fatal("expected false in test environment")
	}
}
