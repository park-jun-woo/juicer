//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindChildByType_Found 테스트
package nestjs

import "testing"

func TestFindChildByType_Found(t *testing.T) {
	src := []byte(`import { X } from './x';`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := findChildByType(root, "import_statement")
	if stmt == nil {
		t.Fatal("expected import_statement")
	}
}
