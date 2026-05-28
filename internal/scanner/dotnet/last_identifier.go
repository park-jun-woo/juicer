//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what qualified_name 노드에서 마지막 identifier를 반환한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func lastIdentifier(node *sitter.Node, src []byte) string {
	last := ""
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == "identifier" {
			last = nodeText(child, src)
		}
	}
	return last
}
