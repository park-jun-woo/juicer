//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 단일 어노테이션에서 HTTP 매핑 메서드와 경로를 매칭한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func matchAnnotationRoute(ann *sitter.Node, src []byte) (string, string, bool) {
	name := annotationName(ann, src)
	if httpMethod, ok := httpMappingAnnotations[name]; ok {
		path := extractAnnotationPath(ann, src)
		return httpMethod, path, true
	}
	if name == AnnRequestMapping {
		path := extractAnnotationPath(ann, src)
		methodStr := annotationElementValue(ann, src, "method")
		httpMethod := resolveRequestMappingMethod(methodStr)
		if httpMethod == "" {
			httpMethod = "GET"
		}
		return httpMethod, path, true
	}
	return "", "", false
}
