//ff:func feature=scan type=test control=sequence
//ff:what TestDetectActix_NoFile 테스트
package scanner

import "testing"

func TestDetectActix_NoFile(t *testing.T) {
	dir := t.TempDir()
	if detectActix(dir) {
		t.Fatal("expected false when Cargo.toml missing")
	}
}
