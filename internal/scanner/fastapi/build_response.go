//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what routeInfo로 scanner.Response를 생성한다
package fastapi

import (
	"strconv"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// buildResponse creates a scanner.Response from route info.
func buildResponse(ri routeInfo, respType string) scanner.Response {
	resp := scanner.Response{Kind: "json"}
	if ri.statusCode > 0 {
		resp.Status = strconv.Itoa(ri.statusCode)
	} else {
		resp.Status = defaultStatusForMethod(ri.method)
	}
	if respType != "" {
		oa := pyTypeToOpenAPI(respType)
		if oa.Type != "" {
			resp.TypeName = respType
		}
	}
	return resp
}
