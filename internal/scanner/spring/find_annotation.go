//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 노드에서 특정 이름의 어노테이션을 찾아 반환한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func findAnnotation(node *sitter.Node, src []byte, name string) *sitter.Node {
	modifiers := findModifiers(node)
	if modifiers == nil {
		return nil
	}
	for _, ann := range childrenOfType(modifiers, "marker_annotation") {
		if annotationName(ann, src) == name {
			return ann
		}
	}
	for _, ann := range childrenOfType(modifiers, "annotation") {
		if annotationName(ann, src) == name {
			return ann
		}
	}
	return nil
}
