//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestCollectHonoVarFromDecl_Multiple 테스트
package hono

import "testing"

func TestCollectHonoVarFromDecl_Multiple(t *testing.T) {
	fi := mustParse(t, []byte(`const app = new Hono(), n = 5;`+"\n"))
	decls := findAllByType(fi.Root, "lexical_declaration")
	vars := map[string]bool{}
	for _, d := range decls {
		collectHonoVarFromDecl(d, fi, vars)
	}
	if !vars["app"] || vars["n"] {
		t.Fatalf("expected only app, got %v", vars)
	}
}
