//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 클래스 레벨 [Authorize] 역할을 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractClassRoles(cls *sitter.Node, src []byte) []string {
	return extractAuthorizeRoles(cls, src)
}
