//ff:func feature=scan type=test control=sequence
//ff:what slicePrefix 테스트
package gogin

import "testing"

func TestSlicePrefix(t *testing.T) {
	if got := slicePrefix(true); got != "[]" {
		t.Errorf("true: got %q", got)
	}
	if got := slicePrefix(false); got != "" {
		t.Errorf("false: got %q", got)
	}
}
