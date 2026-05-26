//ff:func feature=sql type=test control=sequence
//ff:what TestHandleSQLSubcommand 테스트
package main

import "testing"

func TestHandleSQLSubcommand(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()

	if !handleSQLSubcommand([]string{"status"}) {
		t.Fatal("expected true for status")
	}
	if handleSQLSubcommand([]string{"bogus"}) {
		t.Fatal("expected false for unknown subcommand")
	}
}
