//ff:func feature=scan type=test control=sequence
//ff:what newVersionCmd Use 및 RunE 버전 출력 직접 테스트
package main

import "testing"

func TestNewVersionCmd(t *testing.T) {
	cmd := newVersionCmd()
	if cmd.Use != "version" {
		t.Errorf("Use = %q", cmd.Use)
	}
	if err := cmd.RunE(cmd, nil); err != nil {
		t.Errorf("RunE: %v", err)
	}
}
