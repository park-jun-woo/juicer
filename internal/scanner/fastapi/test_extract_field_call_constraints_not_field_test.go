//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractFieldCallConstraints_NotField 테스트
package fastapi

import "testing"

func TestExtractFieldCallConstraints_NotField(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x: int = other(ge=0)\n"))
	f := &pydanticField{}
	extractFieldCallConstraints(assign, src, f)
	if f.ge != nil {
		t.Fatalf("expected no constraints, got %+v", f)
	}
}
