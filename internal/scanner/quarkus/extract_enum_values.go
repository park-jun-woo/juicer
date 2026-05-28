//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what enum 선언에서 상수 값 목록을 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractEnumValues(cls *sitter.Node, src []byte) []string {
	body := findChildByType(cls, "enum_body")
	if body == nil {
		return nil
	}
	var values []string
	for i := 0; i < int(body.ChildCount()); i++ {
		child := body.Child(i)
		if child.Type() != "enum_constant" {
			continue
		}
		nameNode := findChildByType(child, "identifier")
		if nameNode != nil {
			values = append(values, nodeText(nameNode, src))
		}
	}
	return values
}
