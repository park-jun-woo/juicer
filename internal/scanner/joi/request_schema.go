//ff:type feature=scan type=model topic=joi
//ff:what Joi 검증 const(`{ body, query, params }`)에서 추출한 요청 필드 묶음
package joi

import "github.com/park-jun-woo/codistill/internal/scanner"

// RequestSchema — Joi 검증 객체의 body/query/params 필드 묶음
type RequestSchema struct {
	Body   []scanner.Field
	Query  []scanner.Field
	Params []scanner.Field
}

// Empty — 추출된 필드가 하나도 없는지 여부
func (r RequestSchema) Empty() bool {
	return len(r.Body) == 0 && len(r.Query) == 0 && len(r.Params) == 0
}
