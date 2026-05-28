//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 프로퍼티 노드에서 이름을 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractPropertyName(prop *sitter.Node, src []byte) string {
	for i := 0; i < int(prop.ChildCount()); i++ {
		child := prop.Child(i)
		if child.Type() == "identifier" {
			return nodeText(child, src)
		}
	}
	return ""
}
