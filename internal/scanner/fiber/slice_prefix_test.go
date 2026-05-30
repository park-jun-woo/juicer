//ff:func feature=scan type=test control=sequence
//ff:what slicePrefix — 슬라이스 접두 테스트
package fiber

import "testing"

func TestSlicePrefix(t *testing.T) {
	if got := slicePrefix(true); got != "[]" {
		t.Errorf("true -> %q, want []", got)
	}
	if got := slicePrefix(false); got != "" {
		t.Errorf("false -> %q, want empty", got)
	}
}
