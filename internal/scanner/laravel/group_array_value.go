//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 'key' => value 배열 항목에서 => 다음의 값 노드를 반환한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// groupArrayValue returns the value node following the "=>" token in a
// 'key' => value array_element_initializer, or nil when absent.
func groupArrayValue(elem *sitter.Node) *sitter.Node {
	seen := false
	for i := 0; i < int(elem.ChildCount()); i++ {
		child := elem.Child(i)
		if child.Type() == "=>" {
			seen = true
			continue
		}
		if seen {
			return child
		}
	}
	return nil
}
