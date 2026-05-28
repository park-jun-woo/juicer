//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 어노테이션에서 value/name 속성으로 파라미터 이름을 결정한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func resolveAnnotationName(ann *sitter.Node, src []byte, fallback string) string {
	if ann == nil {
		return fallback
	}
	v := firstStringArg(ann, src)
	if v == "" {
		v = annotationElementValue(ann, src, "value")
		if v == "" {
			v = annotationElementValue(ann, src, "name")
		}
	}
	if v != "" {
		return v
	}
	return fallback
}
