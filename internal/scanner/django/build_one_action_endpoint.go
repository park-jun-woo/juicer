//ff:func feature=scan type=extract control=sequence topic=django
//ff:what ViewSet의 단일 @action 엔드포인트를 생성한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildOneActionEndpoint builds a single endpoint for a ViewSet @action method.
func buildOneActionEndpoint(prefix, actionPath, method string, ai actionInfo, vs *viewsetInfo) scanner.Endpoint {
	path := prefix
	if ai.detail {
		path = combinePath(prefix, "{pk}/"+actionPath)
	} else {
		path = combinePath(prefix, actionPath)
	}
	path = ensureLeadingSlash(path)

	ep := scanner.Endpoint{
		Method:  method,
		Path:    path,
		Handler: vs.name + "." + ai.name,
		File:    vs.file,
		Line:    ai.line,
	}
	if ai.detail {
		ep.Request = &scanner.Request{
			PathParams: []scanner.Param{{Name: "pk", Type: "integer"}},
		}
	}
	return ep
}
