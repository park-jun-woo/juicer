//ff:func feature=scan type=test topic=django control=sequence
//ff:what relativeImportSubmodule `from .X import ...` 단일 상대 서브모듈명 추출 테스트
package django

import "testing"

func TestRelativeImportSubmodule(t *testing.T) {
	// single-level relative import
	src := []byte("from .views import urlpatterns\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := findChildByType(root, "import_from_statement")
	if got := relativeImportSubmodule(stmt, src); got != "views" {
		t.Errorf("got %q, want views", got)
	}

	// non-relative import -> ""
	src2 := []byte("from app.views import x\n")
	root2, _ := parsePython(src2)
	stmt2 := findChildByType(root2, "import_from_statement")
	if got := relativeImportSubmodule(stmt2, src2); got != "" {
		t.Errorf("non-relative: got %q", got)
	}
}
