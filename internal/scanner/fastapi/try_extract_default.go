//ff:func feature=scan type=extract control=selection topic=fastapi
//ff:what 단일 자식 노드에서 기본값과 호출 함수명을 추출 시도한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// tryExtractDefault tries to extract default value from a single child node.
// Returns empty strings for nodes that are not values (identifier, type, :, =).
// isNone is true when the default value is Python's None literal.
func tryExtractDefault(child *sitter.Node, src []byte) (defaultVal, defaultCall string, isNone bool) {
	switch child.Type() {
	case "identifier", "type", ":", "=":
		// "identifier" may be "None"
		if child.Type() == "identifier" && nodeText(child, src) == "None" {
			return "", "", true
		}
		return "", "", false
	case "call":
		defaultVal = nodeText(child, src)
		funcNode := findChildByType(child, "identifier")
		if funcNode != nil {
			defaultCall = nodeText(funcNode, src)
		}
		return defaultVal, defaultCall, false
	case "none":
		return "", "", true
	default:
		txt := nodeText(child, src)
		if txt == "None" {
			return "", "", true
		}
		return txt, "", false
	}
}
