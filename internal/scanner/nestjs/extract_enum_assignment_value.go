//ff:func feature=scan type=extract control=selection topic=nestjs
//ff:what enum_body 자식 노드에서 멤버 값을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractEnumAssignmentValue extracts a member value from an enum_body child node.
// For enum_assignment nodes: returns the string value if present, otherwise the key name.
// For property_identifier nodes: returns the identifier text.
// Returns ("", false) for irrelevant node types (commas, braces, etc.).
func extractEnumAssignmentValue(child *sitter.Node, src []byte) (string, bool) {
	switch child.Type() {
	case "enum_assignment":
		strNode := findChildByType(child, "string")
		if strNode != nil {
			return unquoteTS(nodeText(strNode, src)), true
		}
		nameNode := findChildByType(child, "property_identifier")
		if nameNode != nil {
			return nodeText(nameNode, src), true
		}
		return "", false
	case "property_identifier":
		return nodeText(child, src), true
	default:
		return "", false
	}
}
