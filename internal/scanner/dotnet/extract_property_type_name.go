//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 프로퍼티 노드에서 타입명을 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractPropertyTypeName(prop *sitter.Node, src []byte) string {
	for i := 0; i < int(prop.ChildCount()); i++ {
		child := prop.Child(i)
		switch child.Type() {
		case "predefined_type", "generic_name", "nullable_type",
			"array_type", "qualified_name":
			return nodeText(child, src)
		}
	}
	return ""
}
