//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 어노테이션에서 value 속성으로 파라미터 이름을 결정한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func resolveAnnotationName(ann *sitter.Node, src []byte, fallback string) string {
	if ann == nil {
		return fallback
	}
	v := firstStringArg(ann, src)
	if v == "" {
		v = annotationElementValue(ann, src, "value")
	}
	if v != "" {
		return v
	}
	return fallback
}
