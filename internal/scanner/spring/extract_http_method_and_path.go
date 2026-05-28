//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 메서드 어노테이션에서 HTTP 메서드와 경로를 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractHTTPMethodAndPath(m *sitter.Node, src []byte) (string, string, bool) {
	modifiers := findModifiers(m)
	if modifiers == nil {
		return "", "", false
	}
	for _, ann := range childrenOfType(modifiers, "annotation") {
		httpMethod, path, found := matchAnnotationRoute(ann, src)
		if found {
			return httpMethod, path, true
		}
	}
	for _, ann := range childrenOfType(modifiers, "marker_annotation") {
		name := annotationName(ann, src)
		if httpMethod, ok := httpMappingAnnotations[name]; ok {
			return httpMethod, "", true
		}
	}
	return "", "", false
}
