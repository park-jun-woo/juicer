//ff:func feature=sql type=test control=sequence
//ff:what newSQLResetCmd Use/RunE 직접 테스트
package main

import "testing"

func TestNewSQLResetCmd(t *testing.T) {
	cmd := newSQLResetCmd()
	if cmd.Use != "reset" {
		t.Errorf("Use = %q", cmd.Use)
	}
	if cmd.RunE == nil {
		t.Error("RunE nil")
	}
}
