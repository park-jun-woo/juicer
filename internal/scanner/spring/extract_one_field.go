//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 필드 선언 노드에서 scanner.Field를 구성한다
package spring

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractOneField(field *sitter.Node, src []byte) scanner.Field {
	f := scanner.Field{}
	f.Type = extractFieldType(field, src)
	f.Name = extractFieldName(field, src)
	applyJsonProperty(field, src, &f)
	applyValidationAnnotations(field, src, &f)
	applyEmailAnnotation(field, src, &f)
	return f
}
