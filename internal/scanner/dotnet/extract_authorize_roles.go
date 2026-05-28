//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what [Authorize] 어트리뷰트에서 Roles를 추출한다
package dotnet

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func extractAuthorizeRoles(node *sitter.Node, src []byte) []string {
	attr := findAttribute(node, src, AttrAuthorize)
	if attr == nil {
		return nil
	}
	rolesVal := attributeNamedArg(attr, src, "Roles")
	if rolesVal == "" {
		return nil
	}
	var roles []string
	for _, r := range strings.Split(rolesVal, ",") {
		r = strings.TrimSpace(r)
		if r != "" {
			roles = append(roles, r)
		}
	}
	return roles
}
