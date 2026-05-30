//ff:func feature=scan type=test control=sequence topic=hono
//ff:what collectHonoVarFromDecl 테스트
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
