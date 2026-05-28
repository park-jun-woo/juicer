//ff:func feature=scan type=extract control=sequence topic=django
//ff:what ViewSet의 단일 CRUD 메서드 엔드포인트를 생성한다
package django

import "github.com/park-jun-woo/codistill/internal/scanner"

// buildViewSetMethodEndpoint builds a single endpoint for a ViewSet CRUD method.
func buildViewSetMethodEndpoint(prefix string, am actionMethod, vs *viewsetInfo, serializers map[string]serializerInfo) scanner.Endpoint {
	path := prefix
	if am.detail {
		path = combinePath(prefix, "{pk}")
	}
	path = ensureLeadingSlash(path)

	ep := scanner.Endpoint{
		Method:  am.method,
		Path:    path,
		Handler: vs.name + "." + am.action,
		File:    vs.file,
		Line:    vs.line,
	}

	if am.detail {
		ep.Request = &scanner.Request{
			PathParams: []scanner.Param{{Name: "pk", Type: "integer"}},
		}
	}

	addSerializerInfo(&ep, am, vs.serializerClass, serializers)
	return ep
}
