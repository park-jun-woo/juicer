//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveSameFileInterface_Round5 테스트
package spring

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveSameFileInterface_Round5(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "Api.java")
	if err := os.WriteFile(p, []byte(`interface UserApi {}
class UserController implements UserApi {}
`), 0o644); err != nil {
		t.Fatal(err)
	}
	got := resolveSameFileInterface(p, "UserApi")
	if got != p {
		t.Fatalf("expected same file %q, got %q", p, got)
	}
	if got := resolveSameFileInterface(p, "Missing"); got != "" {
		t.Fatalf("expected empty for missing iface, got %q", got)
	}
}
