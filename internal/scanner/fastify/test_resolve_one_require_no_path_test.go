//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveOneRequire_NoPath 테스트
package fastify

import "testing"

func TestResolveOneRequire_NoPath(t *testing.T) {
	d, fi := reqDeclarator(t, "const m = require();\n")
	imports := map[string]string{}
	resolveOneRequire(d, fi.Src, t.TempDir(), imports)
	if len(imports) != 0 {
		t.Fatalf("expected no imports for require with no arg, got %v", imports)
	}
}
