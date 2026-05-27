//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what object 노드에서 "enum" 키의 배열 원소를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractEnumArray finds the "enum" pair in an object node and returns its array elements.
func extractEnumArray(obj *sitter.Node, src []byte) []string {
	for i := 0; i < int(obj.ChildCount()); i++ {
		child := obj.Child(i)
		if child.Type() != "pair" {
			continue
		}
		key := findChildByType(child, "property_identifier")
		if key == nil || nodeText(key, src) != "enum" {
			continue
		}
		arr := findChildByType(child, "array")
		if arr == nil {
			continue
		}
		return collectEnumElements(arr, src)
	}
	return nil
}
