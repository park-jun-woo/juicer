//ff:func feature=sql type=test control=sequence
//ff:what handleSQLSubcommand 전 분기 테스트 (에러 분기는 os.Exit 때문에 서브프로세스 테스트로 분리)
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
