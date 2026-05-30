//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestResolveSamePackageClass 테스트
package quarkus

import (
	"path/filepath"
	"testing"
)

func TestResolveSamePackageClass(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "UserDto.java", `class UserDto {}`)
	referrer := filepath.Join(dir, "R.java")
	if got := resolveSamePackageClass(referrer, "UserDto"); got != filepath.Join(dir, "UserDto.java") {
		t.Fatalf("got %q", got)
	}
	if got := resolveSamePackageClass(referrer, "Missing"); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
