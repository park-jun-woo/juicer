//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what Laravel in: 규칙의 enum 변환 테스트
package laravel

import "testing"

func TestLaravelRulesToField_Enum(t *testing.T) {
	f := laravelRulesToField("role", []string{"required", "in:admin,user,editor"})
	if len(f.Enum) != 3 {
		t.Fatalf("Enum length = %d, want 3", len(f.Enum))
	}
	if f.Enum[0] != "admin" || f.Enum[1] != "user" || f.Enum[2] != "editor" {
		t.Errorf("Enum = %v, want [admin user editor]", f.Enum)
	}
}
