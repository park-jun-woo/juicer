//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what applyFieldConstraint: default/ge/le/min_length/max_length + 음수값 무시 분기
package fastapi

import "testing"

func TestApplyFieldConstraint_Default(t *testing.T) {
	f := &pydanticField{}
	applyFieldConstraint("default", "5", f)
	if !f.hasDefault {
		t.Error("expected hasDefault")
	}
	f2 := &pydanticField{}
	applyFieldConstraint("default_factory", "list", f2)
	if !f2.hasDefault {
		t.Error("expected hasDefault for default_factory")
	}
}

func TestApplyFieldConstraint_NumericConstraints(t *testing.T) {
	f := &pydanticField{}
	applyFieldConstraint("ge", "0", f)
	applyFieldConstraint("le", "100", f)
	applyFieldConstraint("min_length", "1", f)
	applyFieldConstraint("max_length", "20", f)
	if f.ge == nil || *f.ge != 0 {
		t.Errorf("ge=%v", f.ge)
	}
	if f.le == nil || *f.le != 100 {
		t.Errorf("le=%v", f.le)
	}
	if f.minLength == nil || *f.minLength != 1 {
		t.Errorf("minLength=%v", f.minLength)
	}
	if f.maxLength == nil || *f.maxLength != 20 {
		t.Errorf("maxLength=%v", f.maxLength)
	}
}

func TestApplyFieldConstraint_NegativeIgnored(t *testing.T) {
	// non-numeric / negative -> parseIntDefault returns -1 -> not set
	f := &pydanticField{}
	applyFieldConstraint("ge", "abc", f)
	applyFieldConstraint("le", "xyz", f)
	applyFieldConstraint("min_length", "nope", f)
	applyFieldConstraint("max_length", "bad", f)
	if f.ge != nil || f.le != nil || f.minLength != nil || f.maxLength != nil {
		t.Errorf("expected all nil, got %+v", f)
	}
}

func TestApplyFieldConstraint_UnknownKey(t *testing.T) {
	f := &pydanticField{}
	applyFieldConstraint("unknown", "1", f)
	if f.hasDefault || f.ge != nil {
		t.Error("unknown key should be no-op")
	}
}
