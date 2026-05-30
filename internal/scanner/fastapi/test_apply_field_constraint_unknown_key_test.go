//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestApplyFieldConstraint_UnknownKey 테스트
package fastapi

import "testing"

func TestApplyFieldConstraint_UnknownKey(t *testing.T) {
	f := &pydanticField{}
	applyFieldConstraint("unknown", "1", f)
	if f.hasDefault || f.ge != nil {
		t.Error("unknown key should be no-op")
	}
}
