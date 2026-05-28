//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 메서드 본문의 ResponseEntity 패턴에서 HTTP 상태 코드를 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractBodyStatus(m *sitter.Node, src []byte) string {
	body := findChildByType(m, "block")
	if body == nil {
		return ""
	}

	code := matchBodyInvocations(body, src)
	if code != "" {
		return code
	}

	return matchBodyConstructors(body, src)
}
