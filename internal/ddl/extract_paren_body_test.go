package ddl

import "testing"

func TestExtractParenBody_Basic(t *testing.T) {
	got := extractParenBody("CREATE TABLE users (id INT, name TEXT)")
	if got != "id INT, name TEXT" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractParenBody_NoParen(t *testing.T) {
	got := extractParenBody("CREATE TABLE users")
	if got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractParenBody_Nested(t *testing.T) {
	got := extractParenBody("CREATE TABLE t (id INT CHECK(id > 0), name TEXT)")
	if got != "id INT CHECK(id > 0), name TEXT" {
		t.Fatalf("got %q", got)
	}
}
