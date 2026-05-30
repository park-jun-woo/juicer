//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what TestCollectBasePathFromDecl_NoValue 테스트
package hono

import "testing"

func TestCollectBasePathFromDecl_NoValue(t *testing.T) {

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
