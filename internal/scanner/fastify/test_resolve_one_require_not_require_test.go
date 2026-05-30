//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveOneRequire_NotRequire 테스트
package fastify

import "testing"

func TestResolveOneRequire_NotRequire(t *testing.T) {
	d, fi := reqDeclarator(t, `const m = foo("./mod");`+"\n")
	imports := map[string]string{}
	resolveOneRequire(d, fi.Src, t.TempDir(), imports)
	if len(imports) != 0 {
		t.Fatalf("expected no imports for non-require, got %v", imports)
	}
}
