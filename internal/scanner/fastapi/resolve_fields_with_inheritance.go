//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 같은 파일 내 부모 필드를 재귀 병합한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// resolveFieldsWithInheritance returns merged fields (parent + own) for the
// given class. visited prevents infinite recursion on cycles.
func resolveFieldsWithInheritance(cls *sitter.Node, root *sitter.Node, src []byte, visited map[string]bool) []pydanticField {
	nameNode := findChildByType(cls, "identifier")
	if nameNode == nil {
		return extractPydanticFields(cls, src)
	}
	name := nodeText(nameNode, src)
	if visited[name] {
		return nil
	}
	visited[name] = true

	own := extractPydanticFields(cls, src)
	parents := collectParentNames(cls, src)
	var parentFields []pydanticField
	for _, p := range parents {
		if isWellKnown(p) {
			continue
		}
		pf := findParentFieldsInFile(root, src, p, visited)
		parentFields = append(parentFields, pf...)
	}
	return mergeParentAndOwnFields(parentFields, own)
}
