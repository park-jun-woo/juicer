//ff:func feature=sql type=test control=sequence
//ff:what TestHandleSQLSubcommand_All 테스트
package main

import "testing"

func TestHandleSQLSubcommand_All(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()

	if !handleSQLSubcommand([]string{"status"}) {
		t.Fatal("expected true for status")
	}
	if handleSQLSubcommand([]string{"bogus"}) {
		t.Fatal("expected false for unknown subcommand")
	}
	if !handleSQLSubcommand([]string{"next"}) {
		t.Fatal("expected true for next")
	}

	_, cleanup2 := setupSQLSession(t)
	defer cleanup2()
	if !handleSQLSubcommand([]string{"list"}) {
		t.Fatal("expected true for list")
	}
	if !handleSQLSubcommand([]string{"skip"}) {
		t.Fatal("expected true for skip")
	}
	if !handleSQLSubcommand([]string{"reset"}) {
		t.Fatal("expected true for reset")
	}
}
