//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what member_call_expression의 마지막 직계 name 자식(메서드명)을 반환한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func lastMemberCallName(call *sitter.Node, src []byte) string {
	methodName := ""
	for i := 0; i < int(call.ChildCount()); i++ {
		child := call.Child(i)
		if child.Type() == "name" {
			methodName = nodeText(child, src)
		}
	}
	return methodName
}
