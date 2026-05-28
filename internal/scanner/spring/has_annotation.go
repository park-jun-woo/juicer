//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 노드에 특정 어노테이션이 존재하는지 확인한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func hasAnnotation(node *sitter.Node, src []byte, name string) bool {
	modifiers := findModifiers(node)
	if modifiers == nil {
		return false
	}
	for _, ann := range childrenOfType(modifiers, "marker_annotation") {
		if annotationName(ann, src) == name {
			return true
		}
	}
	for _, ann := range childrenOfType(modifiers, "annotation") {
		if annotationName(ann, src) == name {
			return true
		}
	}
	return false
}
