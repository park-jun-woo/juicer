//ff:func feature=scan type=extract control=sequence topic=express
//ff:what new_expression에서 call_expression을 추출하거나 직접 Router 식별자를 확인한다
package express

import sitter "github.com/smacker/go-tree-sitter"

// unwrapNewExpression — new Router() 패턴에서 내부 call_expression을 반환한다.
// identifier만 있으면 원본 node를 반환하여 "매치됨"을 알린다.
// 매치 실패 시 nil을 반환한다.
func unwrapNewExpression(node *sitter.Node, src []byte, aliases map[string]bool) *sitter.Node {
	ce := findChildByType(node, "call_expression")
	if ce != nil {
		return ce
	}
	id := findChildByType(node, "identifier")
	if id != nil && aliases[nodeText(id, src)] {
		return node
	}
	return nil
}
