//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 파라미터 노드에서 타입을 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractParamType(param *sitter.Node, src []byte) string {
	for i := 0; i < int(param.ChildCount()); i++ {
		child := param.Child(i)
		switch child.Type() {
		case "predefined_type", "identifier", "generic_name",
			"nullable_type", "array_type", "qualified_name":
			return nodeText(child, src)
		}
	}
	return ""
}
