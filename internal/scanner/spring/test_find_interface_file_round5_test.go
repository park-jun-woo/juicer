//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestFindInterfaceFile_Round5 테스트
package spring

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindInterfaceFile_Round5(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "UserApi.java"), []byte(`package com.x;
interface UserApi {}
`), 0o644); err != nil {
		t.Fatal(err)
	}
	imports := map[string]string{"UserApi": "com.x.UserApi"}
	got := findInterfaceFile("UserApi", imports, filepath.Join(dir, "UserController.java"), dir)
	if got == "" {
		t.Fatalf("expected to find interface file")
	}
}
