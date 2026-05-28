//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what Laravel 유효성 규칙 → scanner.Field 변환 테스트
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

func TestLaravelRulesToField_Integer(t *testing.T) {
	f := laravelRulesToField("age", []string{"nullable", "integer", "min:0", "max:150"})
	if f.Type != "integer" {
		t.Errorf("Type = %q, want %q", f.Type, "integer")
	}
	if !f.Nullable {
		t.Error("expected Nullable = true")
	}
	if f.Minimum == nil || *f.Minimum != 0 {
		t.Errorf("Minimum = %v, want 0", f.Minimum)
	}
	if f.Maximum == nil || *f.Maximum != 150 {
		t.Errorf("Maximum = %v, want 150", f.Maximum)
	}
}

func TestLaravelRulesToField_Enum(t *testing.T) {
	f := laravelRulesToField("role", []string{"required", "in:admin,user,editor"})
	if len(f.Enum) != 3 {
		t.Fatalf("Enum length = %d, want 3", len(f.Enum))
	}
	if f.Enum[0] != "admin" || f.Enum[1] != "user" || f.Enum[2] != "editor" {
		t.Errorf("Enum = %v, want [admin user editor]", f.Enum)
	}
}

func TestLaravelRulesToField_Email(t *testing.T) {
	f := laravelRulesToField("email", []string{"required", "email"})
	if f.Type != "string" {
		t.Errorf("Type = %q, want %q", f.Type, "string")
	}
}
