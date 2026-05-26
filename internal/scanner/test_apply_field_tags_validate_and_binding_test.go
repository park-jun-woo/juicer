//ff:func feature=scan type=test control=sequence
//ff:what TestApplyFieldTags_ValidateAndBinding 테스트
package scanner

import "testing"

func TestApplyFieldTags_ValidateAndBinding(t *testing.T) {
	f := &Field{Name: "Email"}
	excluded := applyFieldTags(f, `json:"email" validate:"required,email"`)
	if excluded {
		t.Fatal("should not be excluded")
	}
	if f.Validate != "required,email" {
		t.Fatalf("expected required,email, got %s", f.Validate)
	}

	f2 := &Field{Name: "Age"}
	excluded = applyFieldTags(f2, `json:"age" binding:"required"`)
	if excluded {
		t.Fatal("should not be excluded")
	}
	if f2.Validate != "required" {
		t.Fatalf("expected required, got %s", f2.Validate)
	}
}
