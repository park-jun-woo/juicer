//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestCollectHonoVarFromDecl_NonHono 테스트
package hono

import "testing"

func TestCollectHonoVarFromDecl_NonHono(t *testing.T) {
	fi := mustParse(t, []byte("const x = 5;\nlet y;\n"))
	decls := findAllByType(fi.Root, "lexical_declaration")
	vars := map[string]bool{}
	for _, d := range decls {
		collectHonoVarFromDecl(d, fi, vars)
	}
	if len(vars) != 0 {
		t.Fatalf("expected no vars, got %v", vars)
	}
}
