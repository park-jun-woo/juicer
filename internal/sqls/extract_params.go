//ff:func feature=sql type=parse control=iteration dimension=2
//ff:what 함수 파라미터를 "이름 타입" 문자열로 변환
package sqls

import (
	"go/ast"
)

// extractParams formats function parameters as "name type" strings.
//
func extractParams(fields *ast.FieldList) []string {
	if fields == nil {
		return nil
	}
	var params []string
	for _, field := range fields.List {
		typStr := typeString(field.Type)
		if typStr == "context.Context" {
			continue
		}
		if len(field.Names) == 0 {
			params = append(params, typStr)
		} else {
			for _, name := range field.Names {
				params = append(params, name.Name+" "+typStr)
			}
		}
	}
	return params
}

