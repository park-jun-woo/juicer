//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestAppendValidate 테스트
package laravel

import "testing"

func TestAppendValidate(t *testing.T) {
	if got := appendValidate("", "required"); got != "required" {
		t.Fatalf("got %q", got)
	}
	if got := appendValidate("required", "email"); got != "required,email" {
		t.Fatalf("got %q", got)
	}
}
