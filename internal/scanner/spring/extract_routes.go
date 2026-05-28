//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 컨트롤러 메서드에서 HTTP 라우트를 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractMethodEndpoints(cls *sitter.Node, fi *fileInfo) []endpointInfo {
	body := findChildByType(cls, "class_body")
	if body == nil {
		return nil
	}
	methods := childrenOfType(body, "method_declaration")
	var result []endpointInfo
	for _, m := range methods {
		ep, ok := extractOneRoute(m, fi)
		if ok {
			result = append(result, ep)
		}
	}
	return result
}
