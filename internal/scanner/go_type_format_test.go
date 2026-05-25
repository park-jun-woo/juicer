package scanner

import "testing"

func TestGoTypeFormat_Int64(t *testing.T) {
	got := goTypeFormat("int64", Field{})
	if got != "int64" {
		t.Fatalf("expected int64, got %s", got)
	}
}

func TestGoTypeFormat_Float64(t *testing.T) {
	got := goTypeFormat("float64", Field{})
	if got != "double" {
		t.Fatalf("expected double, got %s", got)
	}
}

func TestGoTypeFormat_TimeTime(t *testing.T) {
	got := goTypeFormat("time.Time", Field{})
	if got != "date-time" {
		t.Fatalf("expected date-time, got %s", got)
	}
}

func TestGoTypeFormat_Email(t *testing.T) {
	got := goTypeFormat("string", Field{Validate: "required,email"})
	if got != "email" {
		t.Fatalf("expected email, got %s", got)
	}
}

func TestGoTypeFormat_URI(t *testing.T) {
	got := goTypeFormat("string", Field{Validate: "url"})
	if got != "uri" {
		t.Fatalf("expected uri, got %s", got)
	}
}

func TestGoTypeFormat_Unknown(t *testing.T) {
	got := goTypeFormat("string", Field{})
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
