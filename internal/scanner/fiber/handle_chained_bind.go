//ff:func feature=scan type=extract control=sequence
//ff:what c.Bind().Body(&req) 체이닝 호출에서 요청 바디 타입을 추출한다
package fiber

import (
	"go/ast"
	"go/types"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// handleChainedBind handles c.Bind().Body(req) chained binding pattern.
// outerCall is the .Body(arg) call; its first arg holds the bind target.
func handleChainedBind(ep *scanner.Endpoint, outerCall *ast.CallExpr, method string, info *types.Info) {
	scanner.EnsureRequest(ep)
	if ep.Request.Body != nil {
		return // 첫 번째 바인딩만 기록
	}
	varName := "(unknown)"
	if len(outerCall.Args) > 0 {
		varName = bindVarName(outerCall.Args[0])
	}
	body := &scanner.Body{
		VarName: varName,
		Method:  method,
	}

	typeName, fields := resolveBindType(outerCall, info)
	body.TypeName = typeName
	body.Fields = fields

	ep.Request.Body = body
}
