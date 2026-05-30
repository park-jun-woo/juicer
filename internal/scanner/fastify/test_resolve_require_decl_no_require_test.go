//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestResolveRequireDecl_NoRequire 테스트
package fastify

import "testing"

func TestResolveRequireDecl_NoRequire(t *testing.T) {
	fi := mustParse(t, []byte("const x = 5;\n"))
	decls := findAllByType(fi.Root, "lexical_declaration")
	imports := map[string]string{}
	for _, d := range decls {
		resolveRequireDecl(d, fi.Src, "/d", imports)
	}
	if len(imports) != 0 {
		t.Fatalf("expected no imports, got %v", imports)
	}
}
