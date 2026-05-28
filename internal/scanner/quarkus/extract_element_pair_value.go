//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what element_value_pair 노드에서 지정 키의 값을 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractElementPairValue(child *sitter.Node, src []byte, key string) (string, bool) {
	k := findChildByType(child, "identifier")
	if k == nil {
		return "", false
	}
	if nodeText(k, src) != key {
		return "", false
	}
	foundKey := false
	for j := 0; j < int(child.ChildCount()); j++ {
		val := child.Child(j)
		if val.Type() == "identifier" && !foundKey {
			foundKey = true
			continue
		}
		if val.Type() == "string_literal" {
			return unquoteJava(nodeText(val, src)), true
		}
		if val.Type() == "field_access" {
			return nodeText(val, src), true
		}
		if val.Type() == "identifier" {
			return nodeText(val, src), true
		}
		if val.Type() == "element_value_array_initializer" {
			return nodeText(val, src), true
		}
	}
	return "", false
}
