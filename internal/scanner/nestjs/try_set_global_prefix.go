//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what call_expression이 setGlobalPrefix인지 확인하고 접두사를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// trySetGlobalPrefix checks if a call_expression is setGlobalPrefix and returns the prefix.
func trySetGlobalPrefix(call *sitter.Node, src []byte) (string, bool) {
	memberAccess := findChildByType(call, "member_expression")
	if memberAccess == nil {
		return "", false
	}
	prop := findChildByType(memberAccess, "property_identifier")
	if prop == nil || nodeText(prop, src) != "setGlobalPrefix" {
		return "", false
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return "", false
	}
	return firstStringArg(args, src)
}
