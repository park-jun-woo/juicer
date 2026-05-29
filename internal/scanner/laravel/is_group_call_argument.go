//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what closure가 ->group(...) 호출의 인자로 전달되는지 보고한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// isGroupCallArgument reports whether closure is passed as the argument of a
// ->group(...) call. The closure sits inside argument/arguments wrappers whose
// enclosing member_call_expression names the "group" method.
func isGroupCallArgument(closure *sitter.Node, fi fileInfo) bool {
	for a := closure.Parent(); a != nil; a = a.Parent() {
		switch a.Type() {
		case "argument", "arguments":
			continue
		case "member_call_expression":
			return memberCallName(a, fi.src) == "group"
		default:
			return false
		}
	}
	return false
}
