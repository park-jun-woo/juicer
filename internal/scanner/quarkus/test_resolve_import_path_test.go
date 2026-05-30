//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestResolveImportPath 테스트
package quarkus

import (
	"path/filepath"
	"testing"
)

func TestResolveImportPath(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main/java/com/example/UserDto.java", `class UserDto {}`)
	got := resolveImportPath(dir, "com.example.UserDto")
	if got != filepath.Join(dir, "src/main/java/com/example/UserDto.java") {
		t.Fatalf("got %q", got)
	}
	if got := resolveImportPath(dir, "com.example.Missing"); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
