//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what DTO 필드 목록을 scanner.Field 목록으로 변환한다
package nestjs

import (
	"strings"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// dtoFieldsToScannerFields converts extracted DTO fields to scanner.Field.
func dtoFieldsToScannerFields(fields []dtoField) []scanner.Field {
	var result []scanner.Field
	for _, f := range fields {
		sf := tsTypeToField(f.name, f.tsType, f.optional)
		if len(f.validators) > 0 {
			sf.Validate = strings.Join(f.validators, ",")
		}
		result = append(result, sf)
	}
	return result
}
