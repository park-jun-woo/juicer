//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what import 문의 named imports에서 Router alias 이름을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func collectRouterAliasesFromStmt(stmt *sitter.Node, src []byte, aliases map[string]bool) {
	clause := findChildByType(stmt, "import_clause")
	if clause == nil {
		return
	}
	named := findChildByType(clause, "named_imports")
	if named == nil {
		return
	}
	for _, spec := range childrenOfType(named, "import_specifier") {
		nameNode := spec.ChildByFieldName("name")
		if nameNode == nil {
			continue
		}
		if nodeText(nameNode, src) != "Router" {
			continue
		}
		if alias := spec.ChildByFieldName("alias"); alias != nil {
			aliases[nodeText(alias, src)] = true
		} else {
			aliases["Router"] = true
		}
	}
}
