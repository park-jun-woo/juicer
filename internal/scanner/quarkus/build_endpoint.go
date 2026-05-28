//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what resourceInfo와 endpointInfo로 scanner.Endpoint를 생성한다
package quarkus

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildEndpoint(ri resourceInfo, ep endpointInfo) scanner.Endpoint {
	fullPath := joinPath(ri.prefix, ep.path)

	roles := mergeRoles(ri.roles, ep.roles)

	endpoint := scanner.Endpoint{
		Method:  ep.method,
		Path:    fullPath,
		Handler: ep.handler,
		File:    ep.file,
		Line:    ep.line,
		Roles:   roles,
	}

	req := buildRequest(ep)
	if req != nil {
		endpoint.Request = req
	}
	resp := buildResponse(ep)
	if resp != nil {
		endpoint.Responses = []scanner.Response{*resp}
	}
	return endpoint
}
