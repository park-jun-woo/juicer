//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindChildByType_NotFound 테스트
package nestjs

import "testing"

func TestFindChildByType_NotFound(t *testing.T) {
	src := []byte(`const x = 1;`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := findChildByType(root, "import_statement")
	if stmt != nil {
		t.Fatal("expected nil")
	}
}
