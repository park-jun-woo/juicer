//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 노드에 특정 어트리뷰트가 존재하는지 확인한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func hasAttribute(node *sitter.Node, src []byte, name string) bool {
	return findAttribute(node, src, name) != nil
}
