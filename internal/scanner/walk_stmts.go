//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 문(statement) 목록을 순회하며 라우트·그룹·미들웨어를 수집한다
package scanner

import (
	"go/ast"
	"go/token"
)

func walkStmts(stmts []ast.Stmt, ginAlias, filePath string, fset *token.FileSet, routers map[string]*routerInfo, out *[]Endpoint) {
	for _, stmt := range stmts {
		switch s := stmt.(type) {
		case *ast.AssignStmt:
			processAssign(s, ginAlias, routers)
		case *ast.ExprStmt:
			call, ok := s.X.(*ast.CallExpr)
			if !ok {
				continue
			}
			if ep, ok := tryRouteCall(call, routers, filePath, fset); ok {
				*out = append(*out, ep)
			} else {
				tryUseCall(call, routers)
			}
		case *ast.BlockStmt:
			walkStmts(s.List, ginAlias, filePath, fset, routers, out)
		case *ast.IfStmt:
			if s.Init != nil {
				walkStmts([]ast.Stmt{s.Init}, ginAlias, filePath, fset, routers, out)
			}
			walkStmts(s.Body.List, ginAlias, filePath, fset, routers, out)
			if s.Else != nil {
				walkStmts([]ast.Stmt{s.Else}, ginAlias, filePath, fset, routers, out)
			}
		case *ast.ForStmt:
			walkStmts(s.Body.List, ginAlias, filePath, fset, routers, out)
		case *ast.RangeStmt:
			walkStmts(s.Body.List, ginAlias, filePath, fset, routers, out)
		case *ast.SwitchStmt:
			walkStmts(s.Body.List, ginAlias, filePath, fset, routers, out)
		case *ast.TypeSwitchStmt:
			walkStmts(s.Body.List, ginAlias, filePath, fset, routers, out)
		case *ast.SelectStmt:
			walkStmts(s.Body.List, ginAlias, filePath, fset, routers, out)
		case *ast.CaseClause:
			walkStmts(s.Body, ginAlias, filePath, fset, routers, out)
		case *ast.CommClause:
			walkStmts(s.Body, ginAlias, filePath, fset, routers, out)
		}
	}
}

