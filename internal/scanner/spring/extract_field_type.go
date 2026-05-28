//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 필드 선언에서 타입을 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractFieldType(field *sitter.Node, src []byte) string {
	for i := 0; i < int(field.ChildCount()); i++ {
		child := field.Child(i)
		switch child.Type() {
		case "type_identifier", "generic_type", "array_type", "integral_type",
			"floating_point_type", "boolean_type", "void_type", "scoped_type_identifier":
			return nodeText(child, src)
		}
	}
	return ""
}
