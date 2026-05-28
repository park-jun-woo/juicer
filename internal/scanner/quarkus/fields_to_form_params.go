//ff:func feature=scan type=convert control=iteration dimension=1 topic=quarkus
//ff:what 필드를 폼 파라미터로 변환한다
package quarkus

import "github.com/park-jun-woo/codistill/internal/scanner"

func fieldsToFormParams(fields []scanner.Field) []scanner.Param {
	var params []scanner.Param
	for _, f := range fields {
		name := f.Name
		if f.JSON != "" {
			name = f.JSON
		}
		params = append(params, scanner.Param{
			Name: name,
			Type: f.Type,
		})
	}
	return params
}
