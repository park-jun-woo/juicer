//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 파라미터 노드에서 타입을 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractParamType(param *sitter.Node, src []byte) string {
	for i := 0; i < int(param.ChildCount()); i++ {
		child := param.Child(i)
		switch child.Type() {
		case "type_identifier", "generic_type", "array_type", "integral_type",
			"floating_point_type", "boolean_type", "void_type", "scoped_type_identifier":
			return nodeText(child, src)
		}
	}
	return ""
}
