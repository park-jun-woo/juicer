//ff:func feature=scan type=test control=sequence topic=hono
//ff:what collectBasePathFromDecl 테스트
package hono

import "testing"

func TestCollectBasePathFromDecl_Found(t *testing.T) {
	fi := mustParse(t, []byte(`const app = new Hono().basePath("/api/v1");`+"\n"))
	decls := findAllByType(fi.Root, "lexical_declaration")
	if len(decls) == 0 {
		t.Fatal("no lexical_declaration")
	}
	basePaths := map[string]string{}
	for _, d := range decls {
		collectBasePathFromDecl(d, fi, basePaths)
	}
	if basePaths["app"] != "/api/v1" {
		t.Fatalf("expected /api/v1, got %v", basePaths)
	}
}

func TestCollectBasePathFromDecl_NoBasePath(t *testing.T) {
	fi := mustParse(t, []byte(`const x = 5;`+"\n"))
	decls := findAllByType(fi.Root, "lexical_declaration")
	basePaths := map[string]string{}
	for _, d := range decls {
		collectBasePathFromDecl(d, fi, basePaths)
	}
	if len(basePaths) != 0 {
		t.Fatalf("expected no base paths, got %v", basePaths)
	}
}

func TestCollectBasePathFromDecl_MultipleDeclarators(t *testing.T) {
	fi := mustParse(t, []byte(`const a = new Hono().basePath("/a"), b = 5;`+"\n"))
	decls := findAllByType(fi.Root, "lexical_declaration")
	basePaths := map[string]string{}
	for _, d := range decls {
		collectBasePathFromDecl(d, fi, basePaths)
	}
	if basePaths["a"] != "/a" || len(basePaths) != 1 {
		t.Fatalf("expected only a=/a, got %v", basePaths)
	}
}

func TestCollectBasePathFromDecl_NoValue(t *testing.T) {
	// declarator without a value (let x;) -> value nil -> skipped
	fi := mustParse(t, []byte("let x;\n"))
	decls := findAllByType(fi.Root, "lexical_declaration")
	basePaths := map[string]string{}
	for _, d := range decls {
		collectBasePathFromDecl(d, fi, basePaths)
	}
	if len(basePaths) != 0 {
		t.Fatalf("expected no base paths, got %v", basePaths)
	}
}
