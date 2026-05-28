//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 단일 Endpoint를 생성한다
package express

import "github.com/park-jun-woo/codistill/internal/scanner"

func buildOneEndpoint(method, oaPath string, r routeInfo, relPath string, pathParams []string) scanner.Endpoint {
	ep := scanner.Endpoint{
		Method:     method,
		Path:       oaPath,
		Handler:    r.Handler,
		File:       relPath,
		Line:       r.Line,
		Middleware: r.Middleware,
	}
	if len(pathParams) > 0 {
		var params []scanner.Param
		for _, p := range pathParams {
			params = append(params, scanner.Param{Name: p, Type: "string"})
		}
		ep.Request = &scanner.Request{PathParams: params}
	}
	return ep
}
