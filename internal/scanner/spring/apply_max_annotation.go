//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what @Max 어노테이션에서 최대값을 적용한다
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyMaxAnnotation(field *sitter.Node, src []byte, f *scanner.Field) {
	ann := findAnnotation(field, src, AnnMax)
	if ann == nil {
		return
	}
	if v, ok := singleIntArg(ann, src); ok {
		f.Maximum = intPtr(v)
		return
	}
	if v, ok := annotationIntValue(ann, src, "value"); ok {
		f.Maximum = intPtr(v)
	}
}
