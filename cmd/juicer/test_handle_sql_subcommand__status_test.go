//ff:func feature=sql type=command control=sequence
//ff:what TestHandleSQLSubcommand_Status 테스트
package main

import "testing"

func TestHandleSQLSubcommand_Status(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	got := handleSQLSubcommand([]string{"status"})
	if !got {
		t.Fatal("expected true")
	}
}
