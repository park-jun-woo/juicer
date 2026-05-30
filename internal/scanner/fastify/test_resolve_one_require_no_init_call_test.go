//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveOneRequire_NoInitCall 테스트
package fastify

import "testing"

func TestResolveOneRequire_NoInitCall(t *testing.T) {
	d, fi := reqDeclarator(t, "const m = 5;\n")
	imports := map[string]string{}
	resolveOneRequire(d, fi.Src, t.TempDir(), imports)
	if len(imports) != 0 {
		t.Fatalf("expected no imports, got %v", imports)
	}
}
