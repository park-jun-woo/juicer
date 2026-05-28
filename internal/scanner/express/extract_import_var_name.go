//ff:func feature=scan type=extract control=sequence topic=express
//ff:what import 구문에서 default 또는 named import 변수명을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractImportVarName(stmt *sitter.Node, src []byte) string {
	clause := findChildByType(stmt, "import_clause")
	if clause == nil {
		return ""
	}
	// 1. default import: import x from "..."
	id := findChildByType(clause, "identifier")
	if id != nil {
		return nodeText(id, src)
	}
	// 2. named import: import { x } from "..." or import { x as y } from "..."
	named := findChildByType(clause, "named_imports")
	if named == nil {
		return ""
	}
	spec := findChildByType(named, "import_specifier")
	if spec == nil {
		return ""
	}
	if alias := spec.ChildByFieldName("alias"); alias != nil {
		return nodeText(alias, src)
	}
	if name := spec.ChildByFieldName("name"); name != nil {
		return nodeText(name, src)
	}
	if first := findChildByType(spec, "identifier"); first != nil {
		return nodeText(first, src)
	}
	return ""
}
