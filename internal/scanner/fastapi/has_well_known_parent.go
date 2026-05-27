//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 클래스의 부모 중 well-known이 있는지 확인한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// hasWellKnownParent returns true if the class has at least one
// parent that is a well-known base class.
func hasWellKnownParent(cls *sitter.Node, src []byte) bool {
	parents := collectParentNames(cls, src)
	for _, p := range parents {
		if isWellKnown(p) {
			return true
		}
	}
	return false
}
