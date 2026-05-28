//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 메서드 본문의 Response 패턴에서 HTTP 상태 코드를 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractResponseStatus(m *sitter.Node, src []byte, ep *endpointInfo) {
	body := findChildByType(m, "block")
	if body == nil {
		return
	}
	code := matchResponseInvocations(body, src)
	if code != "" {
		ep.statusCode = code
	}
}
