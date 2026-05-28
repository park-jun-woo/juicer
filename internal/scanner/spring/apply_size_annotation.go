//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what @Size 어노테이션에서 minLength, maxLength를 적용한다
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applySizeAnnotation(field *sitter.Node, src []byte, f *scanner.Field) {
	ann := findAnnotation(field, src, AnnSize)
	if ann == nil {
		return
	}
	if v, ok := annotationIntValue(ann, src, "min"); ok {
		f.MinLength = intPtr(v)
	}
	if v, ok := annotationIntValue(ann, src, "max"); ok {
		f.MaxLength = intPtr(v)
	}
}
