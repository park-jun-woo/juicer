//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractCallStringArg_NoString 테스트
package fastify

import "testing"

func TestExtractCallStringArg_NoString(t *testing.T) {
	fi := mustParse(t, []byte("foo(bar, baz);\n"))
	calls := findAllByType(fi.Root, "call_expression")[0]
	if got := extractCallStringArg(calls, fi.Src); got != "" {
		t.Fatalf("expected empty for no string arg, got %q", got)
	}
}
