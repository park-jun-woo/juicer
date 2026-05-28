//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 메서드 본문의 method_invocation에서 ResponseEntity 팩토리 메서드 상태 코드를 찾는다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func matchBodyInvocations(body *sitter.Node, src []byte) string {
	for _, inv := range findAllByType(body, "method_invocation") {
		code := matchResponseEntityInvocation(inv, src)
		if code != "" {
			return code
		}
	}
	return ""
}
