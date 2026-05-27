//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what scanner.Field 목록을 dtoField 목록으로 역변환한다
package nestjs

import "github.com/park-jun-woo/codistill/internal/scanner"

// scannerFieldsToDTOFields converts scanner.Field slice back to dtoField slice.
func scannerFieldsToDTOFields(fields []scanner.Field) []dtoField {
	var result []dtoField
	for _, f := range fields {
		result = append(result, fieldToDTOField(f))
	}
	return result
}
