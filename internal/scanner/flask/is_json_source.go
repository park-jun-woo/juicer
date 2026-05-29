//ff:func feature=scan type=extract control=selection topic=flask
//ff:what 표현식 노드가 request.json / request.get_json() 접근인지 판별한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// isJSONSource reports whether node is a `request.json` attribute or a
// `request.get_json(...)` call expression.
func isJSONSource(node *sitter.Node, src []byte) bool {
	if node == nil {
		return false
	}
	switch node.Type() {
	case "attribute":
		return nodeText(node, src) == "request.json"
	case "call":
		fn := findChildByType(node, "attribute")
		return fn != nil && nodeText(fn, src) == "request.get_json"
	default:
		return false
	}
}
