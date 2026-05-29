//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what HttpResponse 체인의 .json(arg) 호출에서 첫 인자 노드를 반환한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func findJSONArg(scopedID *sitter.Node, src []byte) *sitter.Node {
	jsonField := findJSONFieldExpr(scopedID, src)
	if jsonField == nil {
		return nil
	}
	callExpr := jsonField.Parent()
	if callExpr == nil || callExpr.Type() != "call_expression" {
		return nil
	}
	args := findChildByType(callExpr, "arguments")
	if args == nil {
		return nil
	}
	return firstArgExpr(args)
}
