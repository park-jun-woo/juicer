//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what endpointInfo로 scanner.Request를 생성한다
package spring

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildRequest(ep endpointInfo) *scanner.Request {
	req := &scanner.Request{
		PathParams: ep.params,
		Query:      ep.query,
		Headers:    ep.headers,
	}
	if len(ep.files) > 0 {
		req.Files = ep.files
	}
	if ep.bodyType != "" {
		req.Body = &scanner.Body{
			VarName:  ep.bodyVarName,
			Method:   "RequestBody",
			TypeName: ep.bodyType,
		}
	}
	if !hasContent(req) {
		return nil
	}
	return req
}
