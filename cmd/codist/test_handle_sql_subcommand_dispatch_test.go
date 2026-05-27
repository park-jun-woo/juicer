//ff:func feature=sql type=test control=sequence
//ff:what TestHandleSQLSubcommand_Dispatch 서브커맨드 디스패치 테스트
package main

import "testing"

func TestHandleSQLSubcommand_Dispatch(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()

	if !handleSQLSubcommand([]string{"status"}) {
		t.Fatal("expected true for status")
	}
	if !handleSQLSubcommand([]string{"list"}) {
		t.Fatal("expected true for list")
	}
	if !handleSQLSubcommand([]string{"skip"}) {
		t.Fatal("expected true for skip")
	}
	if !handleSQLSubcommand([]string{"next"}) {
		t.Fatal("expected true for next")
	}
	if !handleSQLSubcommand([]string{"reset"}) {
		t.Fatal("expected true for reset")
	}
	if handleSQLSubcommand([]string{"unknown"}) {
		t.Fatal("expected false for unknown subcommand")
	}
}
