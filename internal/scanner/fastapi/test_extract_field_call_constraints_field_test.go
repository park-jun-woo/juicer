//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractFieldCallConstraints_Field 테스트
package fastapi

import "testing"

func TestExtractFieldCallConstraints_Field(t *testing.T) {
	assign, src := firstAssignment(t, []byte("age: int = Field(ge=0, le=120)\n"))
	f := &pydanticField{}
	extractFieldCallConstraints(assign, src, f)
	if f.ge == nil || *f.ge != 0 || f.le == nil || *f.le != 120 {
		t.Fatalf("got %+v", f)
	}
}
