//ff:func feature=sql type=parse control=iteration dimension=2
//ff:what 반환 타입을 문자열로 변환
package sqls

import (
	"go/ast"
)

// extractReturns formats return types as strings.
//
func extractReturns(fields *ast.FieldList) []string {
	if fields == nil {
		return nil
	}
	var returns []string
	for _, field := range fields.List {
		typStr := typeString(field.Type)
		if len(field.Names) == 0 {
			returns = append(returns, typStr)
		} else {
			for _, name := range field.Names {
				returns = append(returns, name.Name+" "+typStr)
			}
		}
	}
	return returns
}

