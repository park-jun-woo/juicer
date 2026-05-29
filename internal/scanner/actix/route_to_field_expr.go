//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 인자가 .to(...) 호출이면 그 field_expression을 반환한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func routeToFieldExpr(child *sitter.Node, src []byte) *sitter.Node {
	if child.Type() != "call_expression" {
		return nil
	}
	fe := findChildByType(child, "field_expression")
	if fe == nil {
		return nil
	}
	fid := findChildByType(fe, "field_identifier")
	if fid == nil || nodeText(fid, src) != "to" {
		return nil
	}
	return fe
}
