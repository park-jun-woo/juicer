//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 메서드 반환 타입을 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractReturnInfo(m *sitter.Node, src []byte, ep *endpointInfo) {
	for i := 0; i < int(m.ChildCount()); i++ {
		child := m.Child(i)
		switch child.Type() {
		case "type_identifier", "generic_type", "void_type", "array_type",
			"integral_type", "floating_point_type", "boolean_type", "scoped_type_identifier":
			raw := nodeText(child, src)
			typeName, isArray := unwrapReturnType(raw)
			ep.returnType = typeName
			ep.returnIsArray = isArray
			return
		}
	}
}
