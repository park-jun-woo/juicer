//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what Laravel integer+min/max 규칙 변환 테스트
package laravel

import "testing"

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
