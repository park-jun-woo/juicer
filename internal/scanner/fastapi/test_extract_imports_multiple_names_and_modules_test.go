//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractImports_MultipleNamesAndModules 테스트
package fastapi

import "testing"

func TestExtractImports_MultipleNamesAndModules(t *testing.T) {
	src := []byte("from app.models import User, Item\nfrom .routes import router\n")
	root, _ := parsePython(src)
	imports := extractImports(root, src)
	if len(imports) != 3 {
		t.Fatalf("expected 3, got %d: %+v", len(imports), imports)
	}
}
