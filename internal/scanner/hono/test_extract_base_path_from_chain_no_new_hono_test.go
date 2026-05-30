//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractBasePathFromChain_NoNewHono 테스트
package hono

import "testing"

func TestExtractBasePathFromChain_NoNewHono(t *testing.T) {
	val, src := valueOfDecl(t, `const a = other.basePath("/x");`+"\n")
	if got := extractBasePathFromChain(val, src); got != "" {
		t.Fatalf("expected empty without new Hono, got %q", got)
	}
}
