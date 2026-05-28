//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what endpointInfo로 scanner.Response를 생성한다
package dotnet

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildResponse(ep endpointInfo) *scanner.Response {
	if ep.returnType == "" && ep.statusCode == "" {
		return nil
	}
	resp := scanner.Response{Kind: "json"}
	if ep.statusCode != "" {
		resp.Status = ep.statusCode
	} else {
		resp.Status = defaultStatusForMethod(ep.method)
	}
	if ep.returnType != "" {
		assignReturnTypeInfo(ep, &resp)
	}
	return &resp
}
