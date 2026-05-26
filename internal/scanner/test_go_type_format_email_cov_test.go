//ff:func feature=scan type=test control=sequence
//ff:what TestGoTypeFormat_EmailCov 테스트
package scanner

import "testing"

func TestGoTypeFormat_EmailCov(t *testing.T) {
	f := Field{Validate: "required,email"}
	if got := goTypeFormat("string", f); got != "email" {
		t.Fatalf("expected email, got %s", got)
	}
}
