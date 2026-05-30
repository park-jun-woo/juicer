//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestCollectBasePathFromDecl_MultipleDeclarators 테스트
package hono

import "testing"

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
