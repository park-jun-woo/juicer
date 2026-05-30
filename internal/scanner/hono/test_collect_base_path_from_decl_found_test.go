//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestCollectBasePathFromDecl_Found 테스트
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
