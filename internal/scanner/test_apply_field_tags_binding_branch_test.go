//ff:func feature=scan type=test control=sequence
//ff:what TestApplyFieldTags_BindingBranch 테스트
package scanner

import "testing"

func TestApplyFieldTags_BindingBranch(t *testing.T) {
	f := &Field{Name: "Age"}
	applyFieldTags(f, `json:"age" binding:"required"`)
	if f.Validate != "required" {
		t.Fatalf("expected 'required', got %s", f.Validate)
	}
}
