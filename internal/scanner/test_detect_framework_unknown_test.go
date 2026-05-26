//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFramework_Unknown 알 수 없는 프레임워크 테스트
package scanner

import "testing"

func TestDetectFramework_Unknown(t *testing.T) {
	dir := t.TempDir()
	if DetectFramework(dir) != "" {
		t.Fatal("expected empty")
	}
}
