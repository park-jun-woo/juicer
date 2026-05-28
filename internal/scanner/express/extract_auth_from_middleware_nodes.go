//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 미들웨어 노드를 순회하여 인증 수준과 역할을 분류한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractAuthFromMiddlewareNodes(mwNodes []*sitter.Node, src []byte) (string, []string) {
	authLevel := ""
	var roles []string
	for _, node := range mwNodes {
		name, isCall := extractMiddlewareNameForAuth(node, src)
		if name == "" {
			continue
		}
		cls := classifyAuthMiddleware(name)
		if cls == "auth" || cls == "role" {
			authLevel = "auth_required"
		}
		if cls == "role" && isCall {
			roles = append(roles, extractRoleStrings(node, src)...)
		}
	}
	if authLevel == "" {
		return "public", nil
	}
	return authLevel, roles
}
