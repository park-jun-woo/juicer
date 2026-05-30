//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractFieldCallConstraints_NoCall 테스트
package fastapi

import "testing"

func TestExtractFieldCallConstraints_NoCall(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x: int = 5\n"))
	f := &pydanticField{}
	extractFieldCallConstraints(assign, src, f)
	if f.ge != nil || f.hasDefault {
		t.Fatalf("expected no-op, got %+v", f)
	}

}
