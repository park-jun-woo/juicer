//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 어노테이션에서 value/path 속성의 경로를 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractAnnotationPath(ann *sitter.Node, src []byte) string {
	path := firstStringArg(ann, src)
	if path == "" {
		path = annotationElementValue(ann, src, "value")
		if path == "" {
			path = annotationElementValue(ann, src, "path")
		}
	}
	return path
}
