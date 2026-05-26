//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what call_expression이 setGlobalPrefix 호출인지 확인한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// hasSetGlobalPrefix checks if a call_expression is a setGlobalPrefix call,
// regardless of argument type.
func hasSetGlobalPrefix(call *sitter.Node, src []byte) bool {
	member := findChildByType(call, "member_expression")
	if member == nil {
		return false
	}
	prop := findChildByType(member, "property_identifier")
	return prop != nil && nodeText(prop, src) == "setGlobalPrefix"
}
