//ff:func feature=scan type=extract control=selection topic=fastify
//ff:what object 노드에서 지정 키의 string/identifier/number 값을 반환한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractPairStringOrIdent(obj *sitter.Node, src []byte, keyName string) string {
	val := findPairValue(obj, src, keyName)
	if val == nil {
		return ""
	}
	switch val.Type() {
	case "string":
		return unquoteTS(nodeText(val, src))
	case "identifier":
		return nodeText(val, src)
	case "number":
		return nodeText(val, src)
	default:
		return ""
	}
}
