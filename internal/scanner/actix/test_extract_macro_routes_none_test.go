//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractMacroRoutes_None 테스트
package actix

import "testing"

func TestExtractMacroRoutes_None(t *testing.T) {
	src := `fn plain() {}`
	root, err := parseRust([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	fi := &fileInfo{src: []byte(src), root: root}
	if routes := extractMacroRoutes(fi); len(routes) != 0 {
		t.Fatalf("expected no routes, got %+v", routes)
	}
}
