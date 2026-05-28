//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what schemaInfo의 params/querystring/body를 Request에 적용한다
package fastify

import "github.com/park-jun-woo/codistill/internal/scanner"

func applySchemaToRequest(si *schemaInfo, src []byte, req *scanner.Request, hasReq *bool) {
	if si.Params != nil {
		schemaParams := jsonSchemaToParams(si.Params, src)
		if len(schemaParams) > 0 {
			req.PathParams = schemaParams
			*hasReq = true
		}
	}
	if si.Querystring != nil {
		req.Query = jsonSchemaToParams(si.Querystring, src)
		if len(req.Query) > 0 {
			*hasReq = true
		}
	}
	if si.Body != nil {
		fields := jsonSchemaToFields(si.Body, src)
		if len(fields) > 0 {
			req.Body = &scanner.Body{VarName: "body", Method: "json", Fields: fields}
			*hasReq = true
		}
	}
}
