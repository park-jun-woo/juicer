//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what OmitType 변환을 적용하여 지정 필드를 제외한다
package nestjs

import "github.com/park-jun-woo/juicer/internal/scanner"

// applyOmitType excludes specified fields.
func applyOmitType(fields []scanner.Field, omitNames []string) []dtoField {
	omitSet := make(map[string]struct{}, len(omitNames))
	for _, name := range omitNames {
		omitSet[name] = struct{}{}
	}
	var result []dtoField
	for _, f := range fields {
		if _, omitted := omitSet[f.Name]; omitted {
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
