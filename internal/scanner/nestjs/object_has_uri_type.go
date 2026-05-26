//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 객체 리터럴에서 type: VersioningType.URI 프로퍼티를 확인한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// objectHasURIType checks if { type: VersioningType.URI } is present in an object node.
func objectHasURIType(obj *sitter.Node, src []byte) bool {
	for i := 0; i < int(obj.ChildCount()); i++ {
		child := obj.Child(i)
		if child.Type() != "pair" {
			continue
		}
		key := findChildByType(child, "property_identifier")
		if key == nil || nodeText(key, src) != "type" {
			continue
		}
		val := findChildByType(child, "member_expression")
		if val != nil && nodeText(val, src) == "VersioningType.URI" {
			return true
		}
	}
	return false
}
