//ff:func feature=ratchet type=test control=sequence
//ff:what newSQLNextCmd Use/플래그(repo,queries)/RunE 직접 테스트
package main

import "testing"

func TestNewSQLNextCmd(t *testing.T) {
	cmd := newSQLNextCmd()
	if cmd.Use != "next" {
		t.Errorf("Use = %q", cmd.Use)
	}
	if cmd.Flags().Lookup("repo") == nil || cmd.Flags().Lookup("queries") == nil {
		t.Error("missing repo/queries flag")
	}
}
