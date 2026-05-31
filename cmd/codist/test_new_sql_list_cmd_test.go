//ff:func feature=sql type=test control=sequence
//ff:what newSQLListCmd Use/Args 직접 테스트
package main

import "testing"

func TestNewSQLListCmd(t *testing.T) {
	cmd := newSQLListCmd()
	if cmd.Use != "list" {
		t.Errorf("Use = %q", cmd.Use)
	}
	if cmd.RunE == nil {
		t.Error("RunE nil")
	}
}
