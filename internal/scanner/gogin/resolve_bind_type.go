//ff:func feature=scan type=extract control=sequence
//ff:what 바인딩 호출의 인자에서 변수 타입을 추적하여 TypeName과 Fields를 반환한다
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func resolveBindType(call *ast.CallExpr, info *types.Info) (typeName string, fields []scanner.Field) {
	if len(call.Args) == 0 || info == nil {
		return "", nil
	}

	arg := call.Args[0]

	// &req -> req
	if unary, ok := arg.(*ast.UnaryExpr); ok && unary.Op == token.AND {
		arg = unary.X
	}

	return resolveExprType(arg, info)
}

