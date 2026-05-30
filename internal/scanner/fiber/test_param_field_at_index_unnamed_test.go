//ff:func feature=scan type=test control=sequence
//ff:what TestParamFieldAtIndex_Unnamed 테스트
package fiber

import "testing"

func TestParamFieldAtIndex_Unnamed(t *testing.T) {

	params := paramsOf(t, "package m\nfunc f(int) {}\n")
	_, name := paramFieldAtIndex(params, 0)
	if name != "_" {
		t.Errorf("unnamed idx 0 = %q, want _", name)
	}
}
