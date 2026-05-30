//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestResolveSameFileAndPackageClass 테스트
package spring

import (
	"path/filepath"
	"testing"
)

func TestResolveSameFileAndPackageClass(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "R.java")
	writeFile(t, dir, "R.java", `class R {} class UserDto { String name; }`)
	if resolveSameFileClass(p, "UserDto", dir) != p {
		t.Fatal("same file")
	}
	writeFile(t, dir, "Other.java", `class Other {}`)
	if resolveSamePackageClass(p, "Other") != filepath.Join(dir, "Other.java") {
		t.Fatal("same package")
	}
}
