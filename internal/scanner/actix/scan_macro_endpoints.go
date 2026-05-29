//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what Pass 1: 매크로 라우트에서 엔드포인트를 추출하고 핸들러 인덱스를 채운다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func scanMacroEndpoints(files []*fileInfo, sIdx structIndex, cache map[string][]scanner.Field, handlerFuncs map[string]*handlerInfo) []scanner.Endpoint {
	var endpoints []scanner.Endpoint
	for _, fi := range files {
		for _, mr := range extractMacroRoutes(fi) {
			endpoints = append(endpoints, buildMacroEndpoint(mr, fi, sIdx, cache))
		}
		collectHandlerFuncs(fi, handlerFuncs)
	}
	return endpoints
}
