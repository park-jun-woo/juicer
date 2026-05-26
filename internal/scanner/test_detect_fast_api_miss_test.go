//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFastAPI_Miss 테스트
package scanner

import "testing"

func TestDetectFastAPI_Miss(t *testing.T) {
	dir := t.TempDir()
	if detectFastAPI(dir) {
		t.Fatal("expected false for empty dir")
	}
}
