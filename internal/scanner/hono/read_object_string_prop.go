//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what object 리터럴에서 지정 키의 string 값을 읽는다 (없으면 빈 문자열)
package hono

import sitter "github.com/smacker/go-tree-sitter"

func readObjectStringProp(obj *sitter.Node, key string, src []byte) string {
	for _, pair := range childrenOfType(obj, "pair") {
		keyNode := pair.ChildByFieldName("key")
		if keyNode == nil || nodeText(keyNode, src) != key {
			continue
		}
		valNode := pair.ChildByFieldName("value")
		if valNode != nil && valNode.Type() == "string" {
			return unquoteTS(nodeText(valNode, src))
		}
	}
	return ""
}
