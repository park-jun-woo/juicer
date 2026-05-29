//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what Laravel string+max 규칙 변환 테스트
package laravel

import "testing"

func TestLaravelRulesToField_StringWithMax(t *testing.T) {
	f := laravelRulesToField("name", []string{"required", "string", "max:255"})
	if f.Name != "name" {
		t.Errorf("Name = %q, want %q", f.Name, "name")
	}
	if f.Type != "string" {
		t.Errorf("Type = %q, want %q", f.Type, "string")
	}
	if f.MaxLength == nil || *f.MaxLength != 255 {
		t.Errorf("MaxLength = %v, want 255", f.MaxLength)
	}
}
