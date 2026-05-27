//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 해석된 필드를 엔드포인트의 body 또는 response에 적용한다
package fastapi

import "github.com/park-jun-woo/codistill/internal/scanner"

// applyModelFields applies resolved fields to the appropriate endpoint location.
func applyModelFields(ep *scanner.Endpoint, isBody bool, fields []scanner.Field) {
	if isBody && ep.Request != nil && ep.Request.Body != nil {
		ep.Request.Body.Fields = fields
	}
	if !isBody && len(ep.Responses) > 0 {
		ep.Responses[0].Fields = fields
	}
}
