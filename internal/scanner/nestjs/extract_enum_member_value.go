//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 특정 enum의 멤버명에 해당하는 단일 값을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractEnumMemberValue finds the enum declaration named enumName and returns
// the value of its member named memberName (e.g. enum RouteKey { Asset='assets' }
// with memberName "Asset" returns "assets"). For valueless members the member
// key name is returned. Returns ("", false) when the enum or member is absent.
//
// Unlike collectEnumMemberValues (which returns ALL member values for @IsEnum),
// this resolves a single member by key — needed for member-expression paths
// like @Controller(RouteKey.Asset).
func extractEnumMemberValue(root *sitter.Node, src []byte, enumName, memberName string) (string, bool) {
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
		return lookupEnumBodyMember(body, src, memberName)
	}
	return "", false
}
