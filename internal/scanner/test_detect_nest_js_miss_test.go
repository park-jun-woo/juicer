//ff:func feature=scan type=test control=sequence
//ff:what TestDetectNestJS_Miss 테스트
package scanner

import "testing"

func TestDetectNestJS_Miss(t *testing.T) {
	dir := t.TempDir()
	if detectNestJS(dir) {
		t.Fatal("expected false for empty dir")
	}
}
