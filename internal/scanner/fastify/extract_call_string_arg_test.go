//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what extractCallStringArg 테스트
package fastify

import "testing"

func TestExtractCallStringArg_String(t *testing.T) {
	fi := mustParse(t, []byte(`require("@fastify/autoload");`+"\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Fatal("no call")
	}
	if got := extractCallStringArg(calls[0], fi.Src); got != "@fastify/autoload" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractCallStringArg_FirstStringWins(t *testing.T) {
	fi := mustParse(t, []byte(`foo("first", "second");`+"\n"))
	calls := findAllByType(fi.Root, "call_expression")[0]
	if got := extractCallStringArg(calls, fi.Src); got != "first" {
		t.Fatalf("expected first string, got %q", got)
	}
}

func TestExtractCallStringArg_NoString(t *testing.T) {
	fi := mustParse(t, []byte("foo(bar, baz);\n"))
	calls := findAllByType(fi.Root, "call_expression")[0]
	if got := extractCallStringArg(calls, fi.Src); got != "" {
		t.Fatalf("expected empty for no string arg, got %q", got)
	}
}
