//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what 라우트 하나를 scanner.Endpoint로 구성한다(컨트롤러/요청/응답 해석 포함)
package laravel

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func buildOneEndpoint(absRoot string, ri routeInfo, parsedFiles map[string]*fileInfo) scanner.Endpoint {
	ep := scanner.Endpoint{
		Method:     ri.method,
		Path:       ri.path,
		Handler:    buildHandlerName(ri.controller, ri.action),
		File:       ri.file,
		Line:       ri.line,
		Middleware: ri.middleware,
	}

	cm := resolveRouteController(absRoot, ri, parsedFiles)
	pathParams := applyControllerParamTypes(extractURLParams(ri.path), cm)

	if req := buildRequest(absRoot, pathParams, cm, parsedFiles); req != nil {
		ep.Request = req
	}
	if cm != nil {
		responses := extractResponsesFromMethod(absRoot, cm, parsedFiles)
		if len(responses) > 0 {
			ep.Responses = responses
		}
	}
	return ep
}
