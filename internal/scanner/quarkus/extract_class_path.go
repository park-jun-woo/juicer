//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 클래스의 @Path에서 prefix 경로를 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractClassPath(cls *sitter.Node, src []byte) string {
	ann := findAnnotation(cls, src, AnnPath)
	if ann == nil {
		return ""
	}
	path := firstStringArg(ann, src)
	if path == "" {
		path = annotationElementValue(ann, src, "value")
	}
	return path
}
