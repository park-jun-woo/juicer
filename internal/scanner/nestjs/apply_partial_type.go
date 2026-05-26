//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what PartialType 변환을 적용하여 모든 필드를 optional로 만든다
package nestjs

import "github.com/park-jun-woo/juicer/internal/scanner"

// applyPartialType makes all fields optional.
func applyPartialType(fields []scanner.Field) []dtoField {
	var result []dtoField
	for _, f := range fields {
		df := dtoField{name: f.Name, tsType: f.Type, optional: true}
		if df.tsType == "" {
			df.tsType = "string"
		}
		result = append(result, df)
	}
	return result
}
