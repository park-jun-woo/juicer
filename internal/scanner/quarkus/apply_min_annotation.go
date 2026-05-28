//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what @Min 어노테이션에서 최소값을 적용한다
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyMinAnnotation(field *sitter.Node, src []byte, f *scanner.Field) {
	ann := findAnnotation(field, src, AnnMin)
	if ann == nil {
		return
	}
	if v, ok := singleIntArg(ann, src); ok {
		f.Minimum = intPtr(v)
		return
	}
	if v, ok := annotationIntValue(ann, src, "value"); ok {
		f.Minimum = intPtr(v)
	}
}
