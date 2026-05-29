//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 스키마 노드가 식별자면 TypeBox 변수 맵에서 Type.Object 인자 객체 노드로 치환한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func resolveTypeBoxRef(node *sitter.Node, src []byte, vars map[string]*sitter.Node) *sitter.Node {
	if node == nil || node.Type() != "identifier" || vars == nil {
		return nil
	}
	return vars[nodeText(node, src)]
}
