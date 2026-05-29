//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what field_expression의 수신자(첫 자식) 노드를 반환한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func findFieldReceiver(fieldExpr *sitter.Node) *sitter.Node {
	if fieldExpr.ChildCount() > 0 {
		return fieldExpr.Child(0)
	}
	return nil
}
