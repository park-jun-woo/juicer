//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what call_expression이 new Hono().method() 체인인지 확인한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func isNewHonoChain(node *sitter.Node, src []byte) bool {
	if node.Type() != "call_expression" {
		return false
	}
	return chainContainsNewHono(node, src)
}
