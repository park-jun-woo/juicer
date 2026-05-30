//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractBasePathFromChain_NotBasePath 테스트
package hono

import "testing"

func TestExtractBasePathFromChain_NotBasePath(t *testing.T) {
	val, src := valueOfDecl(t, `const a = new Hono().get("/x", h);`+"\n")
	if got := extractBasePathFromChain(val, src); got != "" {
		t.Fatalf("expected empty for non-basePath, got %q", got)
	}
}
