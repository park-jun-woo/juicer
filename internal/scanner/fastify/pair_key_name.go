//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what pair 노드에서 키 이름을 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func pairKeyName(pair *sitter.Node, src []byte) string {
	key := findChildByType(pair, "property_identifier")
	if key != nil {
		return nodeText(key, src)
	}
	key = findChildByType(pair, "string")
	if key != nil {
		return unquoteTS(nodeText(key, src))
	}
	key = findChildByType(pair, "number")
	if key != nil {
		return nodeText(key, src)
	}
	return ""
}
