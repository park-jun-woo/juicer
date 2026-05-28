//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 함수 body에서 라우터 파라미터를 전달하는 호출을 찾아 재귀적으로 rescan한다
package echo

import (
	"go/ast"
	"go/types"
)

func forwardRouterCalls(stmts []ast.Stmt, paramName, prefix string, parent *routerInfo, info *types.Info, ctx *groupArgCtx, depth int) {
	for _, stmt := range stmts {
		exprStmt, ok := stmt.(*ast.ExprStmt)
		if !ok {
			continue
		}
		call, ok := exprStmt.X.(*ast.CallExpr)
		if !ok {
			continue
		}
		fwdIdx := routerArgIndex(call, paramName)
		if fwdIdx < 0 {
			continue
		}
		fwdCtx := *ctx
		fwdCtx.info = info
		rescanCalleeWithPrefixDepth(call, fwdIdx, prefix, parent, &fwdCtx, depth+1)
	}
}
