//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestApplyFieldConstraint_NegativeIgnored 테스트
package fastapi

import "testing"

func TestApplyFieldConstraint_NegativeIgnored(t *testing.T) {

	f := &pydanticField{}
	applyFieldConstraint("ge", "abc", f)
	applyFieldConstraint("le", "xyz", f)
	applyFieldConstraint("min_length", "nope", f)
	applyFieldConstraint("max_length", "bad", f)
	if f.ge != nil || f.le != nil || f.minLength != nil || f.maxLength != nil {
		t.Errorf("expected all nil, got %+v", f)
	}
}
