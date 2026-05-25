//ff:func feature=scan type=extract control=sequence
//ff:what TestApplyFieldTags_Binding 테스트
package scanner

import "testing"

func TestApplyFieldTags_Binding(t *testing.T) {
	f := &Field{Name: "Email"}
	applyFieldTags(f, `json:"email" binding:"required"`)
	if f.Validate != "required" {
		t.Fatalf("expected validate=required, got %s", f.Validate)
	}
}
