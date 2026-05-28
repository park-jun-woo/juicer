//ff:func feature=scan type=convert control=iteration dimension=1 topic=django
//ff:what Meta.fields 문자열 목록을 scanner.Field 목록으로 변환한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// metaFieldsToScannerFields converts Meta.fields string names to scanner.Field list.
func metaFieldsToScannerFields(names []string) []scanner.Field {
	var fields []scanner.Field
	for _, fname := range names {
		fields = append(fields, scanner.Field{
			Name: fname,
			Type: "string",
			JSON: fname,
		})
	}
	return fields
}
