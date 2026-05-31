//ff:func feature=scan type=test topic=joi control=sequence
//ff:what appendValidate 빈/기존 validate 누적 테스트
package joi

import "testing"

func TestAppendValidate(t *testing.T) {
	if got := appendValidate("", "required"); got != "required" {
		t.Errorf("empty: %q", got)
	}
	if got := appendValidate("required", "email"); got != "required,email" {
		t.Errorf("append: %q", got)
	}
}
