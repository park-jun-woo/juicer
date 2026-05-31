//ff:func feature=scan type=test control=sequence topic=django
//ff:what importedNameMatches dotted_name/aliased_import/기타 분기 직접 테스트
package django

import "testing"

func TestImportedNameMatches(t *testing.T) {
	src := []byte("from .views import urlpatterns as up\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	alias := findAllByType(root, "aliased_import")
	if len(alias) == 0 {
		t.Fatal("no aliased_import")
	}
	if !importedNameMatches(alias[0], "urlpatterns", src) {
		t.Error("aliased urlpatterns should match")
	}
	if importedNameMatches(alias[0], "nope", src) {
		t.Error("aliased nope should not match")
	}

	// plain dotted_name import: the last dotted_name child is the imported name.
	src2 := []byte("from .views import plain\n")
	root2, _ := parsePython(src2)
	stmt2 := findChildByType(root2, "import_from_statement")
	dotted := childrenOfType(stmt2, "dotted_name")
	plain := dotted[len(dotted)-1]
	if !importedNameMatches(plain, "plain", src2) {
		t.Error("plain should match")
	}
}
