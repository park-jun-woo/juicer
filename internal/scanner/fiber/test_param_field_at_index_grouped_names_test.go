//ff:func feature=scan type=test control=sequence
//ff:what TestParamFieldAtIndex_GroupedNames 테스트
package fiber

import "testing"

func TestParamFieldAtIndex_GroupedNames(t *testing.T) {

	params := paramsOf(t, "package m\nfunc f(a, b int) {}\n")
	_, name1 := paramFieldAtIndex(params, 1)
	if name1 != "b" {
		t.Errorf("grouped idx 1 = %q", name1)
	}
}
