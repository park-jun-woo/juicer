package scanner

import "testing"

func TestUnquote_Quoted(t *testing.T) {
	got := unquote(`"hello"`)
	if got != "hello" {
		t.Fatalf("expected hello, got %s", got)
	}
}

func TestUnquote_Backtick(t *testing.T) {
	got := unquote("`hello`")
	if got != "hello" {
		t.Fatalf("expected hello, got %s", got)
	}
}

func TestUnquote_Invalid(t *testing.T) {
	got := unquote("hello")
	if got != "hello" {
		t.Fatalf("expected hello, got %s", got)
	}
}
