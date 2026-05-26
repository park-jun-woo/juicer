//ff:func feature=scan type=test control=sequence
//ff:what TestDetectGoGin_Miss 테스트
package scanner

import "testing"

func TestDetectGoGin_Miss(t *testing.T) {
	dir := t.TempDir()
	if detectGoGin(dir) {
		t.Fatal("expected false for empty dir")
	}
}
