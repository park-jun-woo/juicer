//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 클래스가 BaseModel을 상속하는지 확인한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// isBaseModelSubclass checks if a class inherits from a well-known Pydantic base class.
// It also traces one level of same-file inheritance.
func isBaseModelSubclass(cls *sitter.Node, root *sitter.Node, src []byte) bool {
	parents := collectParentNames(cls, src)
	for _, p := range parents {
		if isWellKnown(p) {
			return true
		}
	}
	// 같은 파일 내 1단계 상속 추적
	for _, p := range parents {
		if isParentWellKnown(root, src, p) {
			return true
		}
	}
	return false
}
