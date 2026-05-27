//ff:func feature=sql type=command control=sequence
//ff:what TestHandleSQLSubcommand_Skip 테스트
package main

import "testing"

func TestHandleSQLSubcommand_Skip(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	got := handleSQLSubcommand([]string{"skip"})
	if !got {
		t.Fatal("expected true")
	}
}
