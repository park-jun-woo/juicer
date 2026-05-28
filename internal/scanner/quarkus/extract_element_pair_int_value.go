//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what element_value_pair 노드에서 지정 키의 int 값을 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractElementPairIntValue(child *sitter.Node, src []byte, key string) (int, bool) {
	k := findChildByType(child, "identifier")
	if k == nil {
		return 0, false
	}
	if nodeText(k, src) != key {
		return 0, false
	}
	for j := 0; j < int(child.ChildCount()); j++ {
		val := child.Child(j)
		if val.Type() == "decimal_integer_literal" {
			return parseInt(nodeText(val, src)), true
		}
	}
	return 0, false
}
