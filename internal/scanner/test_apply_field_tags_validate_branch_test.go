//ff:func feature=scan type=test control=sequence
//ff:what TestApplyFieldTags_ValidateBranch 테스트
package scanner

import "testing"

func TestApplyFieldTags_ValidateBranch(t *testing.T) {
	f := &Field{Name: "Email"}
	applyFieldTags(f, `json:"email" validate:"required,email"`)
	if f.Validate != "required,email" {
		t.Fatalf("expected 'required,email', got %s", f.Validate)
	}
}
