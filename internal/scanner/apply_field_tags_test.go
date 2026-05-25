package scanner

import "testing"

func TestApplyFieldTags_JSON(t *testing.T) {
	f := &Field{Name: "Name"}
	excluded := applyFieldTags(f, `json:"name"`)
	if excluded {
		t.Fatal("should not be excluded")
	}
	if f.JSON != "name" {
		t.Fatalf("expected name, got %s", f.JSON)
	}
}

func TestApplyFieldTags_JSONDash(t *testing.T) {
	f := &Field{Name: "Secret"}
	excluded := applyFieldTags(f, `json:"-"`)
	if !excluded {
		t.Fatal("should be excluded")
	}
}

func TestApplyFieldTags_JSONWithOmitempty(t *testing.T) {
	f := &Field{Name: "Name"}
	excluded := applyFieldTags(f, `json:"name,omitempty"`)
	if excluded {
		t.Fatal("should not be excluded")
	}
	if f.JSON != "name" {
		t.Fatalf("expected name, got %s", f.JSON)
	}
}

func TestApplyFieldTags_Validate(t *testing.T) {
	f := &Field{Name: "Age"}
	applyFieldTags(f, `validate:"required,min=1"`)
	if f.Validate != "required,min=1" {
		t.Fatalf("expected validate tag, got %s", f.Validate)
	}
}

func TestApplyFieldTags_Binding(t *testing.T) {
	f := &Field{Name: "Email"}
	applyFieldTags(f, `json:"email" binding:"required"`)
	if f.Validate != "required" {
		t.Fatalf("expected validate=required, got %s", f.Validate)
	}
}
