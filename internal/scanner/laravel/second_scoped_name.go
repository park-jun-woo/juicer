//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what scoped_call_expression에서 두 번째 name 자식(메서드명)을 반환한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func secondScopedName(call *sitter.Node, src []byte) string {
	foundScope := false
	for i := 0; i < int(call.ChildCount()); i++ {
		child := call.Child(i)
		if child.Type() != "name" {
			continue
		}
		if !foundScope {
			foundScope = true
			continue
		}
		return nodeText(child, src)
	}
	return ""
}
