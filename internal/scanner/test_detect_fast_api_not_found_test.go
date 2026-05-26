//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFastAPI_NotFound 미감지 테스트
package scanner

import "testing"

func TestDetectFastAPI_NotFound(t *testing.T) {
	dir := t.TempDir()
	if detectFastAPI(dir) {
		t.Fatal("expected false")
	}
}
