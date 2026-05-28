//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 노드에서 PreAuthorize, Secured, RolesAllowed 역할을 수집한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractRolesFromNode(node *sitter.Node, src []byte) []string {
	var roles []string
	roles = append(roles, extractPreAuthorizeRoles(node, src)...)
	roles = append(roles, extractSecuredRoles(node, src)...)
	roles = append(roles, extractRolesAllowed(node, src)...)
	return roles
}
