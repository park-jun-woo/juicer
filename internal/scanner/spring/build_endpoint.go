//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what controllerInfo와 endpointInfo로 scanner.Endpoint를 생성한다
package spring

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildEndpoint(ci controllerInfo, ep endpointInfo) scanner.Endpoint {
	fullPath := joinPath(ci.prefix, ep.path)
	fullPath = springPathToOpenAPI(fullPath)

	roles := mergeRoles(ci.roles, ep.roles)

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
