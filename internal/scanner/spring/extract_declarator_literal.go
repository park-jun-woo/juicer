//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what variable_declarator에서 리터럴 값을 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractDeclaratorLiteral(decl *sitter.Node, src []byte) string {
	for j := 0; j < int(decl.ChildCount()); j++ {
		val := decl.Child(j)
		switch val.Type() {
		case "string_literal":
			return unquoteJava(nodeText(val, src))
		case "decimal_integer_literal":
			return nodeText(val, src)
		case "decimal_floating_point_literal":
			return nodeText(val, src)
		case "true", "false":
			return nodeText(val, src)
		}
	}
	return ""
}
