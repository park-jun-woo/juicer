//ff:func feature=scan type=test control=sequence
//ff:what TestDetectNestJS_NotFound package.json 없음 테스트
package scanner

import "testing"

func TestDetectNestJS_NotFound(t *testing.T) {
	dir := t.TempDir()
	if detectNestJS(dir) {
		t.Fatal("expected false")
	}
}
