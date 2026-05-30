//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractCallStringArg_FirstStringWins 테스트
package fastify

import "testing"

func TestExtractCallStringArg_FirstStringWins(t *testing.T) {
	fi := mustParse(t, []byte(`foo("first", "second");`+"\n"))
	calls := findAllByType(fi.Root, "call_expression")[0]
	if got := extractCallStringArg(calls, fi.Src); got != "first" {
		t.Fatalf("expected first string, got %q", got)
	}
}
