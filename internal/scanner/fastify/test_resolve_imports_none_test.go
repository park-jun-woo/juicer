//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveImports_None 테스트
package fastify

import "testing"

func TestResolveImports_None(t *testing.T) {
	fi := mustParse(t, []byte("const x = 1;\n"))
	fi.Path = "/app/main.ts"
	imports := resolveImports(fi, "/app")
	if len(imports) != 0 {
		t.Fatalf("expected no imports, got %v", imports)
	}
}
