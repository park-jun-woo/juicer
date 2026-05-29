//ff:func feature=scan type=extract control=sequence topic=zod
//ff:what AST에서 z.object() 호출 노드를 모두 수집한다
package zod

import sitter "github.com/smacker/go-tree-sitter"

// FindObjectCalls — z.object() 노드 수집
func FindObjectCalls(node *sitter.Node, src []byte) []*sitter.Node {
	var result []*sitter.Node
	walkNodes(node, func(n *sitter.Node) {
		if IsObjectCall(n, src) {
			result = append(result, n)
		}
	})
	return result
}
