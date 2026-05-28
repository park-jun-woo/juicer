//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 클래스 레벨 역할을 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractClassRoles(cls *sitter.Node, src []byte) []string {
	return extractRolesFromNode(cls, src)
}
