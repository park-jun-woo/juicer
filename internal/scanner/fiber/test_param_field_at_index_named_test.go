//ff:func feature=scan type=test control=sequence
//ff:what TestParamFieldAtIndex_Named 테스트
package fiber

import "testing"

func TestParamFieldAtIndex_Named(t *testing.T) {
	params := paramsOf(t, "package m\nfunc f(a int, b string) {}\n")
	_, name0 := paramFieldAtIndex(params, 0)
	if name0 != "a" {
		t.Errorf("idx 0 = %q", name0)
	}
	_, name1 := paramFieldAtIndex(params, 1)
	if name1 != "b" {
		t.Errorf("idx 1 = %q", name1)
	}
}
