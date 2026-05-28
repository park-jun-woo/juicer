//ff:func feature=scan type=convert control=iteration dimension=1 topic=spring
//ff:what 필드의 타입 파라미터를 실제 타입 인자로 치환한다
package spring

import "github.com/park-jun-woo/codistill/internal/scanner"

func substituteTypeParams(fields []scanner.Field, typeParamNames []string, typeArgs []string) []scanner.Field {
	if len(typeParamNames) == 0 || len(typeArgs) == 0 {
		return fields
	}
	paramMap := make(map[string]string)
	for i, name := range typeParamNames {
		if i < len(typeArgs) {
			paramMap[name] = typeArgs[i]
		}
	}
	result := make([]scanner.Field, len(fields))
	copy(result, fields)
	for i := range result {
		result[i].Type = substituteType(result[i].Type, paramMap)
	}
	return result
}
