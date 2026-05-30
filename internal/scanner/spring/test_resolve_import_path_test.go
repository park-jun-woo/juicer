//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveImportPath 테스트
package spring

import (
	"path/filepath"
	"testing"
)

func TestResolveImportPath(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main/java/com/example/UserDto.java", `class UserDto {}`)
	if resolveImportPath(dir, "com.example.UserDto") != filepath.Join(dir, "src/main/java/com/example/UserDto.java") {
		t.Fatal("import path")
	}
	if resolveImportPath(dir, "com.example.Missing") != "" {
		t.Fatal("missing import")
	}
}
