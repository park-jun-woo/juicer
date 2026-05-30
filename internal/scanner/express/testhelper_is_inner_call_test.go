//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what isInnerCall 테스트 헬퍼: 노드가 다른 call_expression의 내부인지 판별
package express

import sitter "github.com/smacker/go-tree-sitter"

// isInnerCall reports whether c is nested inside another call_expression.
func isInnerCall(c *sitter.Node) bool {
	for p := c.Parent(); p != nil; p = p.Parent() {
		if p.Type() == "call_expression" {
			return true
		}
		if p.Type() == "expression_statement" || p.Type() == "program" {
			return false
		}
	}
	return false
}
