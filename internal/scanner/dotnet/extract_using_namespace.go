//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what using 지시문 노드에서 네임스페이스 이름을 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractUsingNamespace(node *sitter.Node, src []byte) string {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		switch child.Type() {
		case "qualified_name", "identifier":
			return nodeText(child, src)
		}
	}
	return ""
}
