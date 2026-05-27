//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what endpointInfo로 scanner.Response를 생성한다
package nestjs

import (
	"strconv"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// buildResponse creates a scanner.Response from endpoint info.
func buildResponse(ep endpointInfo) scanner.Response {
	resp := scanner.Response{Kind: "json"}
	if ep.statusCode > 0 {
		resp.Status = strconv.Itoa(ep.statusCode)
	} else {
		resp.Status = defaultStatusForMethod(ep.method)
	}
	if ep.returnType != "" {
		oa := tsTypeToOpenAPI(ep.returnType)
		if oa.Type != "" {
			resp.TypeName = ep.returnType
		}
	}
	return resp
}
