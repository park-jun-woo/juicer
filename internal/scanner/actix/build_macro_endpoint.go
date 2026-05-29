//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 매크로 라우트 하나를 scanner.Endpoint로 구성한다(path params + extractor + 응답)
package actix

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func buildMacroEndpoint(mr macroRoute, fi *fileInfo, sIdx structIndex, cache map[string][]scanner.Field) scanner.Endpoint {
	ep := scanner.Endpoint{
		Method:  mr.method,
		Path:    mr.path,
		Handler: mr.handler,
		File:    fi.relPath,
	}
	applyPathParams(&ep, mr.path)
	if mr.funcNode != nil {
		applyHandlerSignature(&ep, mr.funcNode, fi.src, sIdx, cache)
	}
	return ep
}
