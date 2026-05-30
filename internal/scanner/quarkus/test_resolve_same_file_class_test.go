//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestResolveSameFileClass 테스트
package quarkus

import (
	"path/filepath"
	"testing"
)

func TestResolveSameFileClass(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "R.java")
	writeFile(t, dir, "R.java", `class R {} class UserDto { String name; }`)
	if got := resolveSameFileClass(p, "UserDto", dir); got != p {
		t.Fatalf("got %q", got)
	}
	if got := resolveSameFileClass(p, "Missing", dir); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
