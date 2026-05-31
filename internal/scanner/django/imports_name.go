//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what import_from_statement이 지정한 이름을 import하는지(별칭 포함) 판별한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// importsName reports whether the import_from_statement imports the given name
// (directly or aliased), e.g. `from .x import urlpatterns as y`.
func importsName(stmt *sitter.Node, name string, src []byte) bool {
	for i := 0; i < int(stmt.ChildCount()); i++ {
		if importedNameMatches(stmt.Child(i), name, src) {
			return true
		}
	}
	return false
}
