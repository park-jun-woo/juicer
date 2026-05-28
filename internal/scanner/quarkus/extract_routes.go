//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 리소스 클래스 메서드에서 HTTP 라우트를 추출한다
package quarkus

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
