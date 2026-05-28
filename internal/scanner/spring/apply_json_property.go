//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what @JsonProperty에서 JSON 필드명을 적용한다
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyJsonProperty(field *sitter.Node, src []byte, f *scanner.Field) {
	ann := findAnnotation(field, src, AnnJsonProperty)
	if ann == nil {
		return
	}
	val := firstStringArg(ann, src)
	if val == "" {
		val = annotationElementValue(ann, src, "value")
	}
	if val != "" {
		f.JSON = val
	}
}
