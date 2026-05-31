//ff:func feature=scan type=test topic=django control=sequence
//ff:what importsName import_from_statement이 이름(직접/별칭)을 임포트하는지 테스트
package django

import "testing"

func TestImportsName(t *testing.T) {
	src := []byte("from .views import urlpatterns as up, other\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := findChildByType(root, "import_from_statement")
	if stmt == nil {
		t.Fatal("no import_from_statement")
	}
	if !importsName(stmt, "urlpatterns", src) {
		t.Error("aliased urlpatterns should match")
	}
	if !importsName(stmt, "other", src) {
		t.Error("direct other should match")
	}
	if importsName(stmt, "missing", src) {
		t.Error("missing should be false")
	}
}
