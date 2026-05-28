//ff:func feature=scan type=convert control=iteration dimension=1 topic=dotnet
//ff:what Field 슬라이스를 Param 슬라이스로 변환한다
package dotnet

import "github.com/park-jun-woo/codistill/internal/scanner"

func fieldsToFormParams(fields []scanner.Field) []scanner.Param {
	var params []scanner.Param
	for _, f := range fields {
		params = append(params, scanner.Param{
			Name: f.Name,
			Type: f.Type,
		})
	}
	return params
}
