//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 노드에서 @RolesAllowed, @Authenticated 역할을 수집한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractRolesFromNode(node *sitter.Node, src []byte) []string {
	var roles []string
	roles = append(roles, extractRolesAllowed(node, src)...)
	if hasAnnotation(node, src, AnnAuthenticated) {
		if len(roles) == 0 {
			roles = append(roles, "**")
		}
	}
	return roles
}
