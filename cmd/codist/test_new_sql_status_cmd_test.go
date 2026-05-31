//ff:func feature=sql type=test control=sequence
//ff:what newSQLStatusCmd Use/RunE 직접 테스트
package main

import "testing"

func TestNewSQLStatusCmd(t *testing.T) {
	cmd := newSQLStatusCmd()
	if cmd.Use != "status" {
		t.Errorf("Use = %q", cmd.Use)
	}
	if cmd.RunE == nil {
		t.Error("RunE nil")
	}
}
