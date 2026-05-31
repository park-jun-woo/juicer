//ff:func feature=sql type=test control=sequence
//ff:what newSQLSkipCmd Use/RunE 직접 테스트
package main

import "testing"

func TestNewSQLSkipCmd(t *testing.T) {
	cmd := newSQLSkipCmd()
	if cmd.Use != "skip" {
		t.Errorf("Use = %q", cmd.Use)
	}
	if cmd.RunE == nil {
		t.Error("RunE nil")
	}
}
