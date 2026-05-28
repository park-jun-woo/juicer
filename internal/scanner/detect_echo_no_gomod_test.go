//ff:func feature=scan type=test control=sequence
//ff:what TestDetectEcho_NoGoMod go.mod 없음 분기 테스트
package scanner

import (
	"testing"
)

func TestDetectEcho_NoGoMod(t *testing.T) {
	dir := t.TempDir()
	if detectEcho(dir) {
		t.Fatal("expected false when go.mod is missing")
	}
}
