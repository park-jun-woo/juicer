//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what member_call_expression의 메서드 이름(첫 name 자식)을 반환한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func memberCallName(call *sitter.Node, src []byte) string {
	for i := 0; i < int(call.ChildCount()); i++ {
		c := call.Child(i)
		if c.Type() == "name" {
			return nodeText(c, src)
		}
	}
	return ""
}
