//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what echo.Context 핸들러 body를 타입정보 없이 스캔하고 AST struct 본문을 보강한다
package echo

import (
	"go/ast"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func scanFallbackFunc(ep *scanner.Endpoint, ft *ast.FuncType, body *ast.BlockStmt, idx *funcIndex) {
	ctxName := echoCtxParamName(ft)
	if ctxName == "" || body == nil {
		return
	}
	before := len(ep.Responses)
	scanBody(ep, body, ctxName, nil, idx, "handler")
	for i := before; i < len(ep.Responses); i++ {
		fillFallbackResponseBody(&ep.Responses[i], idx)
	}
}
