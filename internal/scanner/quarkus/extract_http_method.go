//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 메서드 어노테이션에서 HTTP 메서드를 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractHTTPMethod(m *sitter.Node, src []byte) (string, bool) {
	modifiers := findModifiers(m)
	if modifiers == nil {
		return "", false
	}
	for _, ann := range childrenOfType(modifiers, "marker_annotation") {
		name := annotationName(ann, src)
		if httpMethod, ok := httpMethodAnnotations[name]; ok {
			return httpMethod, true
		}
	}
	for _, ann := range childrenOfType(modifiers, "annotation") {
		name := annotationName(ann, src)
		if httpMethod, ok := httpMethodAnnotations[name]; ok {
			return httpMethod, true
		}
	}
	return "", false
}
