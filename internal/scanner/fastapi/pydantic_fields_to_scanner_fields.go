//ff:func feature=scan type=convert control=iteration dimension=1 topic=fastapi
//ff:what pydanticField 슬라이스를 scanner.Field 슬라이스로 변환한다
package fastapi

import "github.com/park-jun-woo/juicer/internal/scanner"

// pydanticFieldsToScannerFields converts pydantic fields to scanner.Field.
func pydanticFieldsToScannerFields(fields []pydanticField) []scanner.Field {
	if len(fields) == 0 {
		return nil
	}
	result := make([]scanner.Field, len(fields))
	for i, f := range fields {
		result[i] = convertOnePydanticField(f)
	}
	return result
}
