//ff:func feature=scan type=extract control=sequence topic=express
//ff:what Router() 단독 호출(destructure import 패턴)을 인식한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func isRouterStandaloneCall(node *sitter.Node, src []byte, aliases map[string]bool) bool {
	if len(aliases) == 0 {
		return false
	}
	target := node
	if node.Type() == "new_expression" {
		target = unwrapNewExpression(node, src, aliases)
		if target == nil {
			return false
		}
		if target == node {
			return true
		}
	}
	if target.Type() != "call_expression" {
		return false
	}
	fn := findChildByType(target, "identifier")
	if fn == nil {
		return false
	}
	return aliases[nodeText(fn, src)]
}
