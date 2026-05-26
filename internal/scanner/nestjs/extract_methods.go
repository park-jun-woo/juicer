//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 컨트롤러 클래스 본문에서 메서드 목록을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractMethods extracts endpoints from methods in a controller class.
func extractMethods(cls *sitter.Node, src []byte, file string) []endpointInfo {
	body := findChildByType(cls, "class_body")
	if body == nil {
		return nil
	}
	var result []endpointInfo
	methods := childrenOfType(body, "method_definition")
	for _, m := range methods {
		ep, ok := extractOneMethod(m, src, file)
		if ok {
			result = append(result, ep)
		}
	}
	return result
}
