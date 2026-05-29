//ff:func feature=sql type=command control=sequence
//ff:what TestHandleSQLSubcommand_List 테스트
package main

import "testing"

func TestHandleSQLSubcommand_List(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	got := execSQLSub([]string{"list"})
	if !got {
		t.Fatal("expected true")
	}
}
