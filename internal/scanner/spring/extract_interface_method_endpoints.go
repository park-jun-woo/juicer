//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 인터페이스 본문에서 HTTP 메서드 엔드포인트를 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractInterfaceMethodEndpoints(iface *sitter.Node, fi *fileInfo) []endpointInfo {
	body := findChildByType(iface, "interface_body")
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
