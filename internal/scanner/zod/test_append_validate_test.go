//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestAppendValidate 테스트
package zod

import "testing"

func TestAppendValidate(t *testing.T) {
	if appendValidate("", "email") != "email" {
		t.Fatal("empty")
	}
	if appendValidate("required", "email") != "required,email" {
		t.Fatal("append")
	}
}
