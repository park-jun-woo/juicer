//ff:func feature=scan type=test control=sequence
//ff:what TestApplyFieldTags_JSON 테스트
package scanner

import "testing"

func TestApplyFieldTags_JSON(t *testing.T) {
	f := &Field{Name: "Name"}
	excluded := ApplyFieldTags(f, `json:"name"`)
	if excluded {
		t.Fatal("should not be excluded")
	}
	if f.JSON != "name" {
		t.Fatalf("expected name, got %s", f.JSON)
	}

	// validate tag
	f2 := &Field{Name: "Email"}
	ApplyFieldTags(f2, `json:"email" validate:"required,email"`)
	if f2.Validate != "required,email" {
		t.Fatalf("expected validate tag, got %s", f2.Validate)
	}

	// binding tag (when validate is empty)
	f3 := &Field{Name: "Age"}
	ApplyFieldTags(f3, `json:"age" binding:"required"`)
	if f3.Validate != "required" {
		t.Fatalf("expected binding as validate, got %s", f3.Validate)
	}

	// json:"-" should be excluded
	f4 := &Field{Name: "Secret"}
	if !ApplyFieldTags(f4, `json:"-"`) {
		t.Fatal("expected excluded for json:\"-\"")
	}

	// json with comma (omitempty)
	f5 := &Field{Name: "Opt"}
	ApplyFieldTags(f5, `json:"opt,omitempty"`)
	if f5.JSON != "opt" {
		t.Fatalf("expected opt, got %s", f5.JSON)
	}
}
