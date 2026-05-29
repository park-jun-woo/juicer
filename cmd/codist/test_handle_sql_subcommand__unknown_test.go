//ff:func feature=sql type=command control=sequence
//ff:what TestHandleSQLSubcommand_Unknown 테스트
package main

import "testing"

func TestHandleSQLSubcommand_Unknown(t *testing.T) {
	got := execSQLSub([]string{"unknown"})
	if got {
		t.Fatal("expected false for unknown subcommand")
	}
}
