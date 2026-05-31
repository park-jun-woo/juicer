//ff:func feature=scan type=test topic=joi control=sequence
//ff:what joi 테스트 헬퍼 — 주어진 타입의 첫 자손 노드 반환
package joi

import sitter "github.com/smacker/go-tree-sitter"

// firstOfType returns the first descendant node of the given type.
func firstOfType(root *sitter.Node, typ string) *sitter.Node {
	var found *sitter.Node
	walkNodes(root, func(n *sitter.Node) {
		if found == nil && n.Type() == typ {
			found = n
		}
	})
	return found
}
