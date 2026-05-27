//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what PartialType 변환을 적용하여 모든 필드를 optional로 만든다
package nestjs

import "github.com/park-jun-woo/codistill/internal/scanner"

// applyPartialType makes all fields optional.
func applyPartialType(fields []scanner.Field) []dtoField {
	var result []dtoField
	for _, f := range fields {
		df := fieldToDTOField(f)
		df.optional = true
		// validate에서 "required" 제거 → optional로 전환
		df.validate = removeRequired(df.validate)
		result = append(result, df)
	}
	return result
}
