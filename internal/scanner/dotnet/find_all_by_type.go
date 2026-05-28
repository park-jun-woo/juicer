//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what AST 서브트리에서 지정 타입의 모든 노드를 수집한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func findAllByType(node *sitter.Node, typeName string) []*sitter.Node {
	var result []*sitter.Node
	walkNodes(node, func(n *sitter.Node) {
		if n.Type() == typeName {
			result = append(result, n)
		}
	})
	return result
}
