//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what firstOfType 테스트 헬퍼
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

func firstOfType(root *sitter.Node, typ string) *sitter.Node {
	nodes := findAllByType(root, typ)
	if len(nodes) == 0 {
		return nil
	}
	return nodes[0]
}
