//ff:func feature=scan type=extract control=sequence
//ff:what c.BodyParser(&req) 호출에서 변수명과 타입을 추출한다
package fiber

import (
	"go/ast"
	"go/types"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func handleBodyParser(ep *scanner.Endpoint, call *ast.CallExpr, method string, info *types.Info) {
	scanner.EnsureRequest(ep)
	if ep.Request.Body != nil {
		return // 첫 번째 바인딩만 기록
	}
	varName := "(unknown)"
	if len(call.Args) > 0 {
		varName = bindVarName(call.Args[0])
	}
	body := &scanner.Body{
		VarName: varName,
		Method:  method,
	}

	typeName, fields := resolveBindType(call, info)
	body.TypeName = typeName
	body.Fields = fields

	ep.Request.Body = body
}
