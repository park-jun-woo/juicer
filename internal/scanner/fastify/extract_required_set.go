//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what JSON Schema의 required 배열에서 필수 필드명 집합을 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractRequiredSet(schemaNode *sitter.Node, src []byte) map[string]bool {
	reqNode := findPairValue(schemaNode, src, "required")
	if reqNode == nil || reqNode.Type() != "array" {
		return nil
	}
	set := make(map[string]bool)
	for _, elem := range collectArrayElements(reqNode) {
		name := unquoteTS(nodeText(elem, src))
		if name != "" {
			set[name] = true
		}
	}
	return set
}
