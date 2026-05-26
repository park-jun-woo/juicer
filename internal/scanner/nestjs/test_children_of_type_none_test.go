//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestChildrenOfType_None 테스트
package nestjs

import "testing"

func TestChildrenOfType_None(t *testing.T) {
	src := []byte(`const x = 1;`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	stmts := childrenOfType(root, "function_declaration")
	if len(stmts) != 0 {
		t.Fatalf("expected 0, got %d", len(stmts))
	}
}
