//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestCollectHonoVarFromDecl_NewHono 테스트
package hono

import "testing"

func TestCollectHonoVarFromDecl_NewHono(t *testing.T) {
	fi := mustParse(t, []byte(`const app = new Hono();`+"\n"))
	decls := findAllByType(fi.Root, "lexical_declaration")
	vars := map[string]bool{}
	for _, d := range decls {
		collectHonoVarFromDecl(d, fi, vars)
	}
	if !vars["app"] {
		t.Fatalf("expected app, got %v", vars)
	}
}
