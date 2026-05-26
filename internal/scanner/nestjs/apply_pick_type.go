//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what PickType 변환을 적용하여 지정 필드만 포함한다
package nestjs

import "github.com/park-jun-woo/juicer/internal/scanner"

// applyPickType includes only specified fields.
func applyPickType(fields []scanner.Field, pickNames []string) []dtoField {
	pickSet := make(map[string]struct{}, len(pickNames))
	for _, name := range pickNames {
		pickSet[name] = struct{}{}
	}
	var result []dtoField
	for _, f := range fields {
		if _, picked := pickSet[f.Name]; !picked {
			continue
		}
		df := dtoField{name: f.Name, tsType: f.Type}
		if df.tsType == "" {
			df.tsType = "string"
		}
		result = append(result, df)
	}
	return result
}
