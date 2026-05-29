//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what Laravel email 규칙 변환 테스트
package laravel

import "testing"

func TestLaravelRulesToField_Email(t *testing.T) {
	f := laravelRulesToField("email", []string{"required", "email"})
	if f.Type != "string" {
		t.Errorf("Type = %q, want %q", f.Type, "string")
	}
}
