//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractBasePathFromChain_NotCall 테스트
package hono

import "testing"

func TestExtractBasePathFromChain_NotCall(t *testing.T) {
	val, src := valueOfDecl(t, `const x = 5;`+"\n")
	if got := extractBasePathFromChain(val, src); got != "" {
		t.Fatalf("expected empty for non-call, got %q", got)
	}
}
