//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what @DefaultValue에서 기본값을 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractDefaultValue(param *sitter.Node, src []byte) string {
	ann := findAnnotation(param, src, AnnDefaultValue)
	if ann == nil {
		return ""
	}
	return firstStringArg(ann, src)
}
