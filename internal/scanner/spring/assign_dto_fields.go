//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what DTO 필드를 엔드포인트의 request/response에 할당한다
package spring

import "github.com/park-jun-woo/codistill/internal/scanner"

func assignDTOFields(dr dtoRequest, ep *scanner.Endpoint, fields []scanner.Field) {
	if dr.isBody && ep.Request != nil && ep.Request.Body != nil {
		ep.Request.Body.Fields = fields
	}
	if dr.isForm {
		ensureRequest(ep)
		ep.Request.FormFields = fieldsToFormParams(fields)
	}
	if !dr.isBody && !dr.isForm {
		if len(ep.Responses) > 0 {
			ep.Responses[0].Fields = fields
		}
	}
}
