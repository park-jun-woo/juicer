//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestApplyFieldConstraint_Default 테스트
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
