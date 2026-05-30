//ff:func feature=scan type=test control=sequence topic=hono
//ff:what extractFirstStringArg 테스트
package hono

import "testing"

func TestExtractFirstStringArg(t *testing.T) {
	fi := mustParse(t, []byte(`f("/api", x);`+"\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if got := extractFirstStringArg(calls[0], fi.Src); got != "/api" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractFirstStringArg_NoArgs(t *testing.T) {
	fi := mustParse(t, []byte("f();\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if got := extractFirstStringArg(calls[0], fi.Src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestExtractFirstStringArg_TemplateNotString(t *testing.T) {
	// template string is not "string" type -> ""
	fi := mustParse(t, []byte("f(`/api`);\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if got := extractFirstStringArg(calls[0], fi.Src); got != "" {
		t.Fatalf("expected empty for template string, got %q", got)
	}
}

func TestExtractFirstStringArg_NotString(t *testing.T) {
	fi := mustParse(t, []byte("f(x, y);\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if got := extractFirstStringArg(calls[0], fi.Src); got != "" {
		t.Fatalf("expected empty for non-string, got %q", got)
	}
}
