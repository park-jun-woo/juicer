//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 클래스에서 static final 필드의 리터럴 값을 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func findStaticFinalField(cls *sitter.Node, src []byte, fieldName string) string {
	body := findChildByType(cls, "class_body")
	if body == nil {
		return ""
	}
	for i := 0; i < int(body.ChildCount()); i++ {
		child := body.Child(i)
		if child.Type() != "field_declaration" {
			continue
		}
		if !hasModifiers(child, src, "static", "final") {
			continue
		}
		decl := findChildByType(child, "variable_declarator")
		if decl == nil {
			continue
		}
		nameNode := findChildByType(decl, "identifier")
		if nameNode == nil || nodeText(nameNode, src) != fieldName {
			continue
		}
		return extractDeclaratorLiteral(decl, src)
	}
	return ""
}
