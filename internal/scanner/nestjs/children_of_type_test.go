//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestChildrenOfType_Found 테스트
package nestjs

import "testing"

func TestChildrenOfType_Found(t *testing.T) {
	src := []byte(`const x = 1; const y = 2;`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	stmts := childrenOfType(root, "lexical_declaration")
	if len(stmts) != 2 {
		t.Fatalf("expected 2, got %d", len(stmts))
	}
}
