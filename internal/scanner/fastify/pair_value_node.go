//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what pair 노드에서 콜론 뒤의 값 노드를 반환한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func pairValueNode(pair *sitter.Node) *sitter.Node {
	foundColon := false
	for i := 0; i < int(pair.ChildCount()); i++ {
		child := pair.Child(i)
		if child.Type() == ":" {
			foundColon = true
			continue
		}
		if foundColon {
			return child
		}
	}
	return nil
}
