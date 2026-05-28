//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 메서드 본문의 method_invocation에서 Response 상태 코드를 찾는다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func matchResponseInvocations(body *sitter.Node, src []byte) string {
	for _, inv := range findAllByType(body, "method_invocation") {
		code := matchResponseInvocation(inv, src)
		if code != "" {
			return code
		}
	}
	return ""
}
