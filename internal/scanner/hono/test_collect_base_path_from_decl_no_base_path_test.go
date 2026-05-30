//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestCollectBasePathFromDecl_NoBasePath 테스트
package hono

import "testing"

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
