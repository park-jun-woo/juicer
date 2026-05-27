//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what scanner.Field 목록을 쿼리 파라미터 목록으로 변환한다
package nestjs

import "github.com/park-jun-woo/codistill/internal/scanner"

// fieldsToQueryParams converts scanner.Field slice to scanner.Param slice for query parameters.
func fieldsToQueryParams(fields []scanner.Field) []scanner.Param {
	params := make([]scanner.Param, 0, len(fields))
	for _, f := range fields {
		p := scanner.Param{Name: f.Name, Type: f.Type}
		if p.Type == "" {
			p.Type = "string"
		}
		params = append(params, p)
	}
	return params
}
