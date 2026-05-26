//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what endpointInfo와 컨트롤러 정보로 scanner.Endpoint를 생성한다
package nestjs

import (
	"strings"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// buildEndpoint creates a scanner.Endpoint from controller + endpoint info.
func buildEndpoint(globalPrefix string, uriVersioning bool, ci controllerInfo, ep endpointInfo) scanner.Endpoint {
	versionPrefix := ""
	if uriVersioning && ci.version != "" {
		versionPrefix = "v" + ci.version
	}
	fullPath := joinParts(globalPrefix, versionPrefix, ci.prefix, ep.path)
	fullPath = pathToOpenAPI(fullPath)
	if !strings.HasPrefix(fullPath, "/") {
		fullPath = "/" + fullPath
	}
	endpoint := scanner.Endpoint{
		Method:     ep.method,
		Path:       fullPath,
		Handler:    ep.handler,
		File:       ep.file,
		Line:       ep.line,
		Middleware: ep.middleware,
		Roles:      ep.roles,
	}
	req := buildRequest(ep)
	if req != nil {
		endpoint.Request = req
	}
	if ep.statusCode > 0 || ep.returnType != "" {
		endpoint.Responses = []scanner.Response{buildResponse(ep)}
	}
	return endpoint
}
