//ff:func feature=scan type=extract control=selection topic=django
//ff:what 단일 키워드 제약조건을 필드에 적용한다
package django

import (
	sitter "github.com/smacker/go-tree-sitter"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// applyOneConstraint applies a single keyword constraint to a field.
func applyOneConstraint(field *scanner.Field, key string, child *sitter.Node, src []byte) {
	switch key {
	case "max_length":
		field.MaxLength = extractIntValue(child, src)
	case "min_length":
		field.MinLength = extractIntValue(child, src)
	case "max_value":
		field.Maximum = extractIntValue(child, src)
	case "min_value":
		field.Minimum = extractIntValue(child, src)
	case "allow_null":
		if hasTrue(child, src) {
			field.Nullable = true
		}
	case "choices":
		field.Enum = extractChoiceValues(child, src)
	}
}
