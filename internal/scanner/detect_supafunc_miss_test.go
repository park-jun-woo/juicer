//ff:func feature=scan type=test control=sequence
//ff:what TestDetectSupaFunc_Miss 빈 디렉토리에서 감지 실패 테스트
package scanner

import (
	"testing"
)

func TestDetectSupaFunc_Miss(t *testing.T) {
	dir := t.TempDir()
	if detectSupaFunc(dir) {
		t.Fatal("expected false for empty dir")
	}
}
