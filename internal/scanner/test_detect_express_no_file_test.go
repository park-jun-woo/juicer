//ff:func feature=scan type=test control=sequence
//ff:what TestDetectExpress_NoFile 테스트
package scanner

import "testing"

func TestDetectExpress_NoFile(t *testing.T) {
	if detectExpress(t.TempDir()) {
		t.Fatal("expected false when package.json missing")
	}
}
