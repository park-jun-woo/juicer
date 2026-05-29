//ff:func feature=sql type=test control=sequence
//ff:what TestHandleSQLSubcommand_Dispatch 서브커맨드 디스패치 테스트
package main

import "testing"

func TestHandleSQLSubcommand_Dispatch(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()

	if !execSQLSub([]string{"status"}) {
		t.Fatal("expected true for status")
	}
	if !execSQLSub([]string{"list"}) {
		t.Fatal("expected true for list")
	}
	if !execSQLSub([]string{"skip"}) {
		t.Fatal("expected true for skip")
	}
	if !execSQLSub([]string{"next"}) {
		t.Fatal("expected true for next")
	}
	if !execSQLSub([]string{"reset"}) {
		t.Fatal("expected true for reset")
	}
	if execSQLSub([]string{"unknown"}) {
		t.Fatal("expected false for unknown subcommand")
	}
}
