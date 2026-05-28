//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what object의 pair 자식 중 지정 키의 문자열 값을 반환한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractPairStringValue(obj *sitter.Node, src []byte, keyName string) string {
	for _, pair := range childrenOfType(obj, "pair") {
		key := findChildByType(pair, "property_identifier")
		if key == nil || nodeText(key, src) != keyName {
			continue
		}
		val := findChildByType(pair, "string")
		if val != nil {
			return unquoteTS(nodeText(val, src))
		}
	}
	return ""
}
