//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 문(statement) 목록을 순회하며 라우트·그룹·미들웨어를 수집한다
package gogin

import (
	"go/ast"
	"go/token"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func walkStmts(stmts []ast.Stmt, ginAlias, filePath string, fset *token.FileSet, routers map[string]*routerInfo, out *[]scanner.Endpoint, hmap map[int][]ast.Expr) {
	for _, stmt := range stmts {
		switch s := stmt.(type) {
		case *ast.AssignStmt:
			processAssign(s, ginAlias, routers)
		case *ast.ExprStmt:
			call, ok := s.X.(*ast.CallExpr)
			if !ok {
				continue
			}
			if ep, exprs, ok := tryRouteCall(call, routers, filePath, fset); ok {
				hmap[len(*out)] = exprs
				*out = append(*out, ep)
			} else {
				tryUseCall(call, routers)
			}
		case *ast.BlockStmt:
			walkStmts(s.List, ginAlias, filePath, fset, routers, out, hmap)
		case *ast.IfStmt:
			if s.Init != nil {
				walkStmts([]ast.Stmt{s.Init}, ginAlias, filePath, fset, routers, out, hmap)
			}
			walkStmts(s.Body.List, ginAlias, filePath, fset, routers, out, hmap)
			if s.Else != nil {
				walkStmts([]ast.Stmt{s.Else}, ginAlias, filePath, fset, routers, out, hmap)
			}
		case *ast.ForStmt:
			walkStmts(s.Body.List, ginAlias, filePath, fset, routers, out, hmap)
		case *ast.RangeStmt:
			walkStmts(s.Body.List, ginAlias, filePath, fset, routers, out, hmap)
		case *ast.SwitchStmt:
			walkStmts(s.Body.List, ginAlias, filePath, fset, routers, out, hmap)
		case *ast.TypeSwitchStmt:
			walkStmts(s.Body.List, ginAlias, filePath, fset, routers, out, hmap)
		case *ast.SelectStmt:
			walkStmts(s.Body.List, ginAlias, filePath, fset, routers, out, hmap)
		case *ast.CaseClause:
			walkStmts(s.Body, ginAlias, filePath, fset, routers, out, hmap)
		case *ast.CommClause:
			walkStmts(s.Body, ginAlias, filePath, fset, routers, out, hmap)
		}
	}
}
