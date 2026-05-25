//ff:func feature=scan type=extract control=sequence
//ff:what TestApplyFieldTags_Validate 테스트
package scanner

import "testing"

func TestApplyFieldTags_Validate(t *testing.T) {
	f := &Field{Name: "Age"}
	applyFieldTags(f, `validate:"required,min=1"`)
	if f.Validate != "required,min=1" {
		t.Fatalf("expected validate tag, got %s", f.Validate)
	}
}
