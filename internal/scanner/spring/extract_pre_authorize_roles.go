//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what @PreAuthorize에서 역할을 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractPreAuthorizeRoles(node *sitter.Node, src []byte) []string {
	ann := findAnnotation(node, src, AnnPreAuthorize)
	if ann == nil {
		return nil
	}
	expr := firstStringArg(ann, src)
	if expr == "" {
		return nil
	}
	var roles []string
	for _, m := range hasRoleRegexp.FindAllStringSubmatch(expr, -1) {
		if len(m) > 1 {
			roles = append(roles, normalizeRole(m[1]))
		}
	}
	for _, m := range hasAnyRoleRegexp.FindAllStringSubmatch(expr, -1) {
		if len(m) > 1 {
			roles = append(roles, extractRolesFromExpr(m[1])...)
		}
	}
	return roles
}
