//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractBasePathFromChain_Match 테스트
package hono

import "testing"

func TestExtractBasePathFromChain_Match(t *testing.T) {
	val, src := valueOfDecl(t, `const a = new Hono().basePath("/api");`+"\n")
	if got := extractBasePathFromChain(val, src); got != "/api" {
		t.Fatalf("got %q, want /api", got)
	}
}
