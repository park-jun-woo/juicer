package scanner

import "testing"

func TestIsRequired_True(t *testing.T) {
	if !isRequired(Field{Validate: "required"}) {
		t.Fatal("expected true")
	}
}

func TestIsRequired_InComma(t *testing.T) {
	if !isRequired(Field{Validate: "required,email"}) {
		t.Fatal("expected true")
	}
}

func TestIsRequired_False(t *testing.T) {
	if isRequired(Field{Validate: "email"}) {
		t.Fatal("expected false")
	}
}

func TestIsRequired_Empty(t *testing.T) {
	if isRequired(Field{}) {
		t.Fatal("expected false for empty validate")
	}
}
