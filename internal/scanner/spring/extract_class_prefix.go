//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 클래스의 @RequestMapping에서 prefix 경로를 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractClassPrefix(cls *sitter.Node, src []byte) string {
	ann := findAnnotation(cls, src, AnnRequestMapping)
	if ann == nil {
		return ""
	}
	path := firstStringArg(ann, src)
	if path == "" {
		path = annotationElementValue(ann, src, "value")
		if path == "" {
			path = annotationElementValue(ann, src, "path")
		}
	}
	return path
}
