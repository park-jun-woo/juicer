//ff:func feature=ddl type=test control=sequence
//ff:what newDDLCmd Use/플래그 및 RunE 기본 디렉터리 실행 직접 테스트
package main

import (
	"strings"
	"testing"
)

func TestNewDDLCmd(t *testing.T) {
	cmd := newDDLCmd()
	if !strings.HasPrefix(cmd.Use, "ddl") {
		t.Errorf("Use = %q", cmd.Use)
	}
	if cmd.Flags().Lookup("output") == nil {
		t.Error("missing --output flag")
	}
	// RunE with no args defaults dir to "." (empty dir -> no error)
	dir := t.TempDir()
	cmd.SetArgs([]string{dir})
	if err := cmd.Execute(); err != nil {
		t.Errorf("execute on empty dir: %v", err)
	}
}
