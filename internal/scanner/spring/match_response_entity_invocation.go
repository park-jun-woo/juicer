//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what method_invocation 노드에서 ResponseEntity 팩토리 메서드 패턴을 매칭한다
package spring

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func matchResponseEntityInvocation(inv *sitter.Node, src []byte) string {
	text := nodeText(inv, src)

	if strings.Contains(text, "ResponseEntity.ok(") {
		return "200"
	}
	if strings.Contains(text, "ResponseEntity.created(") {
		return "201"
	}
	if strings.Contains(text, "ResponseEntity.noContent()") {
		return "204"
	}
	if strings.Contains(text, "ResponseEntity.badRequest()") {
		return "400"
	}
	if strings.Contains(text, "ResponseEntity.notFound()") {
		return "404"
	}
	if strings.Contains(text, "ResponseEntity.status(") {
		return extractStatusArg(inv, src)
	}

	return ""
}
