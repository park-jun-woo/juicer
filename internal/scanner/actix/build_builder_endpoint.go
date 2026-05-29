//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 빌더 라우트 하나를 scanner.Endpoint로 구성하고 핸들러 시그니처를 해석한다
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func buildBuilderEndpoint(br builderRoute, sIdx structIndex, cache map[string][]scanner.Field, handlerFuncs map[string]*handlerInfo) scanner.Endpoint {
	ep := scanner.Endpoint{
		Method:  br.method,
		Path:    br.path,
		Handler: br.handler,
	}
	applyPathParams(&ep, br.path)
	if hi, ok := handlerFuncs[br.handler]; ok {
		ep.File = hi.file.relPath
		applyHandlerSignature(&ep, hi.funcNode, hi.file.src, sIdx, cache)
	}
	return ep
}
