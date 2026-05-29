//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what AST struct 필드 하나를 (이름별로) scanner.Field로 변환해 누적한다
package fiber

import (
	"go/ast"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func appendAstField(fields []scanner.Field, field *ast.Field) []scanner.Field {
	if len(field.Names) == 0 {
		return fields
	}
	typeStr := exprString(field.Type)
	jsonName := astFieldJSONName(field)
	for _, name := range field.Names {
		if !name.IsExported() {
			continue
		}
		fields = append(fields, scanner.Field{Name: name.Name, Type: typeStr, JSON: jsonName})
	}
	return fields
}
