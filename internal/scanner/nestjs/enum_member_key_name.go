//ff:func feature=scan type=extract control=selection topic=nestjs
//ff:what enum_body 자식 노드의 멤버 키 이름을 반환한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// enumMemberKeyName returns the member key name (the identifier on the left of
// "=", or the bare identifier for valueless members) of an enum_body child.
func enumMemberKeyName(child *sitter.Node, src []byte) string {
	switch child.Type() {
	case "enum_assignment":
		nameNode := findChildByType(child, "property_identifier")
		if nameNode != nil {
			return nodeText(nameNode, src)
		}
	case "property_identifier":
		return nodeText(child, src)
	}
	return ""
}
