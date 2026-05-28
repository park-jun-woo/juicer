//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what import 구문에서 default 또는 named import 변수명을 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractImportVarName(stmt *sitter.Node, src []byte) string {
	clause := findChildByType(stmt, "import_clause")
	if clause == nil {
		return ""
	}
	id := findChildByType(clause, "identifier")
	if id != nil {
		return nodeText(id, src)
	}
	return extractNamedImportVarName(clause, src)
}
