//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what fn() -> Scope 정의의 본문에서 반환식인 web::scope(...) 체인 call_expression을 찾는다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// funcReturnScopeChain returns the call_expression in the function body whose
// receiver chain head is web::scope (the `web::scope(...).service(...)` builder
// chain that a `fn xxx_scope() -> Scope` returns). It scans the function block
// and picks the first such chain so the caller can feed it back through the
// web::scope branch of processServiceArg for prefix synthesis.
func funcReturnScopeChain(funcNode *sitter.Node, src []byte) *sitter.Node {
	block := findChildByType(funcNode, "block")
	if block == nil {
		return nil
	}
	var found *sitter.Node
	walkNodes(block, func(n *sitter.Node) {
		if found != nil || n.Type() != "call_expression" {
			return
		}
		if receiverIsWebScopeOrResource(n, src) {
			found = n
		}
	})
	return found
}
