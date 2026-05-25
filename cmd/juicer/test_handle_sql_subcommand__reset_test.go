//ff:func feature=sql type=command control=sequence
//ff:what TestHandleSQLSubcommand_Reset 테스트
package main

import "testing"

func TestHandleSQLSubcommand_Reset(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	got := handleSQLSubcommand([]string{"reset"})
	if !got {
		t.Fatal("expected true")
	}
}
