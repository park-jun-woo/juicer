//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 노드가 web::scope/web::resource 체인이 아닌 최상위 .service() 호출인지 판별한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func isTopLevelServiceCall(n *sitter.Node, src []byte) bool {
	if n.Type() != "call_expression" {
		return false
	}
	fe := findChildByType(n, "field_expression")
	if fe == nil {
		return false
	}
	fid := findChildByType(fe, "field_identifier")
	if fid == nil {
		return false
	}
	if nodeText(fid, src) != "service" {
		return false
	}
	receiver := findFieldReceiver(fe)
	if receiver != nil && isWebScopeOrResource(receiver, src) {
		return false
	}
	return true
}
