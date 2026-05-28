//ff:func feature=scan type=test control=sequence
//ff:what TestDetectSpring_NotFound — 빈 디렉토리에서 false 반환
package scanner

import "testing"

func TestDetectSpring_NotFound(t *testing.T) {
	dir := t.TempDir()
	if detectSpring(dir) {
		t.Error("expected false for empty dir")
	}
}
