//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 어트리뷰트 노드에서 이름을 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func attributeName(attr *sitter.Node, src []byte) string {
	for i := 0; i < int(attr.ChildCount()); i++ {
		child := attr.Child(i)
		switch child.Type() {
		case "identifier":
			return nodeText(child, src)
		case "qualified_name":
			return lastIdentifier(child, src)
		}
	}
	return ""
}
