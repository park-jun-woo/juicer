//ff:func feature=scan type=extract control=sequence
//ff:what 타입 필드와 태그에서 Field 구조를 생성한다 (json:"-" 필드는 nil 반환)
package echo

import (
	"go/types"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

// buildField creates a Field from a struct field and tag. Returns nil if the field should be excluded (json:"-").
func buildField(f *types.Var, tag string, visited map[string]bool) *scanner.Field {
	field := scanner.Field{
		Name: f.Name(),
		Type: formatType(f.Type()),
	}

	if tag != "" {
		if excluded := scanner.ApplyFieldTags(&field, tag); excluded {
			return nil
		}
	}

	field.Fields = resolveNestedFields(f.Type(), visited)

	return &field
}
