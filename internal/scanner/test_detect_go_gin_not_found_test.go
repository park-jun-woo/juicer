//ff:func feature=scan type=test control=sequence
//ff:what TestDetectGoGin_NotFound go.mod 없음 테스트
package scanner

import "testing"

func TestDetectGoGin_NotFound(t *testing.T) {
	dir := t.TempDir()
	if detectGoGin(dir) {
		t.Fatal("expected false for missing go.mod")
	}
}
