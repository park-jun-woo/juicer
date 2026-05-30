//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestCollectHonoVarFromDecl_Chain 테스트
package hono

import "testing"

func TestCollectHonoVarFromDecl_Chain(t *testing.T) {
	fi := mustParse(t, []byte(`const api = new Hono().basePath("/api");`+"\n"))
	decls := findAllByType(fi.Root, "lexical_declaration")
	vars := map[string]bool{}
	for _, d := range decls {
		collectHonoVarFromDecl(d, fi, vars)
	}
	if !vars["api"] {
		t.Fatalf("expected api from chain, got %v", vars)
	}
}
