//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what TypeScript enum 정의에서 멤버 값을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractEnumMembers finds the enum declaration with the given name and returns its member values.
// For members with string values (e.g. OPEN = 'open'), the string value is returned.
// For members without values (e.g. Up, Down), the key name is returned.
func extractEnumMembers(root *sitter.Node, src []byte, enumName string) []string {
	enums := findAllByType(root, "enum_declaration")
	for _, e := range enums {
		nameNode := findChildByType(e, "identifier")
		if nameNode == nil || nodeText(nameNode, src) != enumName {
			continue
		}
		body := findChildByType(e, "enum_body")
		if body == nil {
			continue
		}
		return collectEnumMemberValues(body, src)
	}
	return nil
}
