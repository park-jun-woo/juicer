//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what endpointInfo로 scanner.Request를 생성한다
package dotnet

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildRequest(ep endpointInfo) *scanner.Request {
	req := &scanner.Request{
		PathParams: ep.params,
		Query:      ep.query,
		Headers:    ep.headers,
		FormFields: ep.formFields,
	}
	if len(ep.files) > 0 {
		req.Files = ep.files
	}
	if ep.bodyType != "" {
		req.Body = &scanner.Body{
			VarName:  ep.bodyVarName,
			Method:   "FromBody",
			TypeName: ep.bodyType,
		}
	}
	if !hasContent(req) {
		return nil
	}
	return req
}
