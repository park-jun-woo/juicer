//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFiber_NoGoMod go.mod가 없는 경우 테스트
package scanner

import (
	"testing"
)

func TestDetectFiber_NoGoMod(t *testing.T) {
	dir := t.TempDir()
	if detectFiber(dir) {
		t.Fatal("expected false")
	}
}
