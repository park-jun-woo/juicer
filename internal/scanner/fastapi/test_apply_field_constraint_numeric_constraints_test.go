//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestApplyFieldConstraint_NumericConstraints 테스트
package fastapi

import "testing"

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
