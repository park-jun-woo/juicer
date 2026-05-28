//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 파라미터 노드에서 이름을 추출한다 (identifier, required_parameter, optional_parameter)
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func paramNodeName(child *sitter.Node, src []byte) string {
	if child.Type() == "identifier" {
		return nodeText(child, src)
	}
	if child.Type() == "required_parameter" || child.Type() == "optional_parameter" {
		id := findChildByType(child, "identifier")
		if id != nil {
			return nodeText(id, src)
		}
	}
	return ""
}
