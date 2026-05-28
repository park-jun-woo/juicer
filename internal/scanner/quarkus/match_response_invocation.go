//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what method_invocation 노드에서 Response 팩토리 메서드 패턴을 매칭한다
package quarkus

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func matchResponseInvocation(inv *sitter.Node, src []byte) string {
	text := nodeText(inv, src)
	for pattern, code := range responseStatusMethods {
		if strings.Contains(text, pattern+"(") || strings.Contains(text, pattern+".") {
			return code
		}
	}
	if strings.Contains(text, "Response.status(") {
		return extractResponseStatusArg(inv, src)
	}
	return ""
}
