//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what import 문에서 named import 이름들을 수집한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func collectImportedNames(stmt *sitter.Node, src []byte) []string {
	clause := findChildByType(stmt, "import_clause")
	if clause == nil {
		return nil
	}
	var names []string
	if id := findChildByType(clause, "identifier"); id != nil {
		names = append(names, nodeText(id, src))
	}
	named := findChildByType(clause, "named_imports")
	if named == nil {
		return names
	}
	for _, spec := range childrenOfType(named, "import_specifier") {
		names = append(names, extractSpecImportName(spec, src))
	}
	return names
}
