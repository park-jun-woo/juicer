//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 문 목록을 순회하며 routers 맵을 구축하고, Group 인자를 포함한 함수 호출을 감지한다
package echo

import (
	"go/ast"
)

func walkForGroupArgs(stmts []ast.Stmt, ctx *groupArgCtx) {
	for _, stmt := range stmts {
		switch s := stmt.(type) {
		case *ast.AssignStmt:
			processAssign(s, ctx.echoAlias, ctx.routers)
		case *ast.ExprStmt:
			call, ok := s.X.(*ast.CallExpr)
			if !ok {
				continue
			}
			tryUseCall(call, ctx.routers)
			tryGroupArgCall(call, ctx)
		case *ast.BlockStmt:
			walkForGroupArgs(s.List, ctx)
		case *ast.IfStmt:
			if s.Init != nil {
				walkForGroupArgs([]ast.Stmt{s.Init}, ctx)
			}
			walkForGroupArgs(s.Body.List, ctx)
			if s.Else != nil {
				walkForGroupArgs([]ast.Stmt{s.Else}, ctx)
			}
		case *ast.ForStmt:
			walkForGroupArgs(s.Body.List, ctx)
		case *ast.RangeStmt:
			walkForGroupArgs(s.Body.List, ctx)
		case *ast.SwitchStmt:
			walkForGroupArgs(s.Body.List, ctx)
		case *ast.TypeSwitchStmt:
			walkForGroupArgs(s.Body.List, ctx)
		case *ast.SelectStmt:
			walkForGroupArgs(s.Body.List, ctx)
		case *ast.CaseClause:
			walkForGroupArgs(s.Body, ctx)
		case *ast.CommClause:
			walkForGroupArgs(s.Body, ctx)
		}
	}
}
