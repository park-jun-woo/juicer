//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what Pass 2: 빌더 라우트에서 엔드포인트를 추출한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func scanBuilderEndpoints(files []*fileInfo, sIdx structIndex, cache map[string][]scanner.Field, handlerFuncs map[string]*handlerInfo) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for _, fi := range files {
		for _, br := range extractBuilderRoutes(fi, handlerFuncs) {
			endpoints = append(endpoints, buildBuilderEndpoint(br, sIdx, cache, handlerFuncs))
		}
	}
	return endpoints
}
