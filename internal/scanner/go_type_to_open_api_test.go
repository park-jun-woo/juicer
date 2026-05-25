package scanner

import "testing"

func TestGoTypeToOpenAPI_String(t *testing.T) {
	if goTypeToOpenAPI("string") != "string" {
		t.Fatal("expected string")
	}
}

func TestGoTypeToOpenAPI_Int(t *testing.T) {
	if goTypeToOpenAPI("int") != "integer" {
		t.Fatal("expected integer")
	}
}

func TestGoTypeToOpenAPI_Float(t *testing.T) {
	if goTypeToOpenAPI("float64") != "number" {
		t.Fatal("expected number")
	}
}

func TestGoTypeToOpenAPI_Bool(t *testing.T) {
	if goTypeToOpenAPI("bool") != "boolean" {
		t.Fatal("expected boolean")
	}
}

func TestGoTypeToOpenAPI_Pointer(t *testing.T) {
	if goTypeToOpenAPI("*string") != "string" {
		t.Fatal("expected string")
	}
}

func TestGoTypeToOpenAPI_Slice(t *testing.T) {
	if goTypeToOpenAPI("[]int") != "array" {
		t.Fatal("expected array")
	}
}

func TestGoTypeToOpenAPI_TimeTime(t *testing.T) {
	if goTypeToOpenAPI("time.Time") != "string" {
		t.Fatal("expected string")
	}
}

func TestGoTypeToOpenAPI_Any(t *testing.T) {
	if goTypeToOpenAPI("any") != "object" {
		t.Fatal("expected object")
	}
}

func TestGoTypeToOpenAPI_Unknown(t *testing.T) {
	if goTypeToOpenAPI("CustomType") != "object" {
		t.Fatal("expected object")
	}
}
