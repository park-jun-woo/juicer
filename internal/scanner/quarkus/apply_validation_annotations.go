//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what @NotNull, @Min, @Max, @Size 유효성 검증 어노테이션을 적용한다
package quarkus

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func applyValidationAnnotations(field *sitter.Node, src []byte, f *scanner.Field) {
	if hasAnnotation(field, src, AnnNotNull) || hasAnnotation(field, src, AnnNotBlank) || hasAnnotation(field, src, AnnNotEmpty) {
		f.Validate = "required"
	}
	applyMinAnnotation(field, src, f)
	applyMaxAnnotation(field, src, f)
	applySizeAnnotation(field, src, f)
}
