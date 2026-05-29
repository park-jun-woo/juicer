//ff:func feature=sql type=command control=sequence
//ff:what TestHandleSQLSubcommand_Next 테스트
package main

import "testing"

func TestHandleSQLSubcommand_Next(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	got := execSQLSub([]string{"next"})
	if !got {
		t.Fatal("expected true")
	}
}
