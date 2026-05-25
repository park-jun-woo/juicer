//ff:func feature=scan type=extract control=sequence
//ff:what TestGoTypeFormat_Email 테스트
package scanner

import "testing"

func TestGoTypeFormat_Email(t *testing.T) {
	got := goTypeFormat("string", Field{Validate: "required,email"})
	if got != "email" {
		t.Fatalf("expected email, got %s", got)
	}
}
