//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractFirstStringArg_NotString 테스트
package hono

import "testing"

func TestExtractFirstStringArg_NotString(t *testing.T) {
	fi := mustParse(t, []byte("f(x, y);\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if got := extractFirstStringArg(calls[0], fi.Src); got != "" {
		t.Fatalf("expected empty for non-string, got %q", got)
	}
}
