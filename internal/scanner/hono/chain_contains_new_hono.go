//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what call_expression 체인에 new Hono() 호출이 포함되어 있는지 확인한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func chainContainsNewHono(node *sitter.Node, src []byte) bool {
	newExprs := findAllByType(node, "new_expression")
	for _, ne := range newExprs {
		if isNewHonoCall(ne, src) {
			return true
		}
	}
	return false
}
