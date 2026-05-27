//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what routeInfo로 scanner.Request를 생성한다
package fastapi

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildRequest creates a scanner.Request from route info.
func buildRequest(ri routeInfo) *scanner.Request {
	req := &scanner.Request{
		PathParams: ri.params,
		Query:      ri.query,
	}
	if len(ri.files) > 0 {
		req.Files = ri.files
	}
	if ri.bodyType != "" {
		req.Body = &scanner.Body{
			VarName:  ri.bodyVarName,
			Method:   "Body",
			TypeName: ri.bodyType,
			Alias:    ri.bodyAlias,
			Embed:    ri.bodyEmbed,
		}
	}
	if !hasContent(req) {
		return nil
	}
	return req
}
