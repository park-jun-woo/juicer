//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 메서드 체인을 수신자 방향으로만 내려가 체인 헤드의 호출 함수명을 반환한다(인자 서브트리는 보지 않음)
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// chainHeadCallRoot walks DOWN the receiver chain (field_expression Child(0))
// only, never descending into call arguments, and returns the scoped/plain
// function name at the chain head. Unlike findCallRoot it ignores web::scope/
// web::resource that appear inside arguments of an earlier link, so it can tell
// `cfg.service(A)` (head "cfg") apart from `web::scope("/p").service(A)`
// (head "web::scope").
func chainHeadCallRoot(node *sitter.Node, src []byte) string {
	for n := node; n != nil; {
		if n.Type() != "call_expression" {
			return nodeText(n, src)
		}
		fn := n.Child(0)
		if fn == nil {
			return ""
		}
		if fn.Type() == "scoped_identifier" || fn.Type() == "identifier" {
			return nodeText(fn, src)
		}
		if fn.Type() != "field_expression" {
			return ""
		}
		n = findFieldReceiver(fn)
	}
	return ""
}
