//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractBasePathFromChain_NoMemberExpr 테스트
package hono

import "testing"

func TestExtractBasePathFromChain_NoMemberExpr(t *testing.T) {

	val, src := valueOfDecl(t, `const a = plainCall("/x");`+"\n")
	if got := extractBasePathFromChain(val, src); got != "" {
		t.Fatalf("expected empty for plain call, got %q", got)
	}
}
