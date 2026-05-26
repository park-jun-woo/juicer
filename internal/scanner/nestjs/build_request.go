//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what endpointInfo로 scanner.Request를 생성한다
package nestjs

import "github.com/park-jun-woo/juicer/internal/scanner"

// buildRequest creates a scanner.Request from endpoint info.
func buildRequest(ep endpointInfo) *scanner.Request {
	req := &scanner.Request{
		PathParams: ep.params,
		Query:      ep.query,
	}
	if len(ep.files) > 0 {
		req.Files = ep.files
	}
	if ep.bodyType != "" {
		req.Body = &scanner.Body{
			Method:   "Body",
			TypeName: ep.bodyType,
		}
	}
	if !hasContent(req) {
		return nil
	}
	return req
}
