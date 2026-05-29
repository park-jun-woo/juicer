//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 메서드 호출 체인을 수신자 방향으로 내려가며 각 .<method>() 호출 인자에 fn을 적용한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// walkMethodChain descends a method-call chain from its outermost call toward
// the chain head, invoking fn(arguments) for each .<method>() call on the way.
// Pattern: web::scope("...").service(X).service(Y) nests with the outermost
// call (.service(Y)) wrapping its receiver (.service(X)), down to the head
// (web::scope("...")). Walking DOWN the receiver chain — rather than up through
// ancestors — keeps the traversal inside this chain and avoids re-entering the
// enclosing .service() that wraps it (which caused unbounded recursion).
func walkMethodChain(node *sitter.Node, src []byte, method string, fn func(*sitter.Node)) {
	for n := node; n != nil && n.Type() == "call_expression"; {
		fe := findChildByType(n, "field_expression")
		if fe == nil {
			return
		}
		invokeIfMethodCall(n, fe, src, method, fn)
		n = fe.Child(0)
	}
}
