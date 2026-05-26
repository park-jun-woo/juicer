//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what call_expression이 enableVersioning(URI)인지 확인한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// isEnableURIVersioning checks if a call_expression is
// app.enableVersioning({ type: VersioningType.URI }).
func isEnableURIVersioning(call *sitter.Node, src []byte) bool {
	member := findChildByType(call, "member_expression")
	if member == nil {
		return false
	}
	prop := findChildByType(member, "property_identifier")
	if prop == nil || nodeText(prop, src) != "enableVersioning" {
		return false
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return false
	}
	obj := findChildByType(args, "object")
	if obj == nil {
		// enableVersioning() with no args defaults to URI
		return true
	}
	return objectHasURIType(obj, src)
}
