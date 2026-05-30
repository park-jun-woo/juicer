//ff:func feature=scan type=test control=sequence
//ff:what TestParamFieldAtIndex_OutOfRange 테스트
package fiber

import "testing"

func TestParamFieldAtIndex_OutOfRange(t *testing.T) {
	params := paramsOf(t, "package m\nfunc f(a int) {}\n")
	field, name := paramFieldAtIndex(params, 5)
	if field != nil || name != "" {
		t.Errorf("out of range should be nil,'', got %v %q", field, name)
	}
}
