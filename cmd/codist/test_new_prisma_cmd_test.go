//ff:func feature=prisma type=test control=sequence
//ff:what newPrismaCmd Use/플래그 및 RunE 실행 직접 테스트
package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewPrismaCmd(t *testing.T) {
	cmd := newPrismaCmd()
	if !strings.HasPrefix(cmd.Use, "prisma") {
		t.Errorf("Use = %q", cmd.Use)
	}
	if cmd.Flags().Lookup("output") == nil {
		t.Error("missing --output flag")
	}
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "schema.prisma"), []byte("model A { id Int @id }"), 0o644); err != nil {
		t.Fatal(err)
	}
	cmd.SetArgs([]string{dir})
	if err := cmd.Execute(); err != nil {
		t.Errorf("execute: %v", err)
	}
}
