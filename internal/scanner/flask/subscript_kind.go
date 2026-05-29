//ff:func feature=scan type=extract control=selection topic=flask
//ff:what subscript 노드의 베이스를 form/json 으로 분류한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// subscriptKind classifies the base of a subscript node as "form", "json", or
// "" (unknown). jsonVars holds local variables bound to a JSON source.
func subscriptKind(node *sitter.Node, src []byte, jsonVars map[string]bool) string {
	base := node.Child(0)
	if base == nil {
		return ""
	}
	switch base.Type() {
	case "attribute":
		return attributeBaseKind(base, src)
	case "call":
		if isJSONSource(base, src) {
			return "json"
		}
	case "identifier":
		if jsonVars[nodeText(base, src)] {
			return "json"
		}
	}
	return ""
}
