//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractFieldConstraints 테스트
package fastapi

import "testing"

func TestExtractFieldConstraints(t *testing.T) {

	args, src := firstArgList(t, []byte("x = Field(..., ge=1, le=9, min_length=2, max_length=8)\n"))
	f := &pydanticField{}
	extractFieldConstraints(args, src, f)
	if f.ge == nil || *f.ge != 1 || f.le == nil || *f.le != 9 {
		t.Fatalf("ge/le wrong: %+v", f)
	}
	if f.minLength == nil || *f.minLength != 2 || f.maxLength == nil || *f.maxLength != 8 {
		t.Fatalf("length wrong: %+v", f)
	}

}
