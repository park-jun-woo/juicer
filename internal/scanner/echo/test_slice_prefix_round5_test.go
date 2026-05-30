//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestSlicePrefix_Round5 테스트
package echo

import "testing"

func TestSlicePrefix_Round5(t *testing.T) {
	if slicePrefix(true) != "[]" {
		t.Fatal("slice")
	}
	if slicePrefix(false) != "" {
		t.Fatal("non-slice")
	}
}
