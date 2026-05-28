//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 필드 선언에서 변수명을 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractFieldName(field *sitter.Node, src []byte) string {
	decl := findChildByType(field, "variable_declarator")
	if decl == nil {
		return ""
	}
	nameNode := findChildByType(decl, "identifier")
	if nameNode == nil {
		return ""
	}
	return nodeText(nameNode, src)
}
