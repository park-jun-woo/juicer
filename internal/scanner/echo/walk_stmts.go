//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 문(statement) 목록을 순회하며 라우트·그룹·미들웨어를 수집한다
package echo

import (
	"go/ast"
	"go/token"
	"go/types"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func walkStmts(info *types.Info, stmts []ast.Stmt, echoAlias, filePath string, fset *token.FileSet, routers map[string]*routerInfo, out *[]scanner.Endpoint, hmap map[int][]ast.Expr) {
	for _, stmt := range stmts {
		switch s := stmt.(type) {
		case *ast.AssignStmt:
			processAssign(info, s, echoAlias, routers)
		case *ast.ExprStmt:
			call, ok := s.X.(*ast.CallExpr)
			if !ok {
				continue
			}
			if ep, exprs, ok := tryRouteCall(info, call, routers, filePath, fset); ok {
				hmap[len(*out)] = exprs
				*out = append(*out, ep)
			} else {
				tryUseCall(call, routers)
			}
		case *ast.BlockStmt:
			walkStmts(info, s.List, echoAlias, filePath, fset, routers, out, hmap)
		case *ast.IfStmt:
			if s.Init != nil {
				walkStmts(info, []ast.Stmt{s.Init}, echoAlias, filePath, fset, routers, out, hmap)
			}
			walkStmts(info, s.Body.List, echoAlias, filePath, fset, routers, out, hmap)
			if s.Else != nil {
				walkStmts(info, []ast.Stmt{s.Else}, echoAlias, filePath, fset, routers, out, hmap)
			}
		case *ast.ForStmt:
			walkStmts(info, s.Body.List, echoAlias, filePath, fset, routers, out, hmap)
		case *ast.RangeStmt:
			walkStmts(info, s.Body.List, echoAlias, filePath, fset, routers, out, hmap)
		case *ast.SwitchStmt:
			walkStmts(info, s.Body.List, echoAlias, filePath, fset, routers, out, hmap)
		case *ast.TypeSwitchStmt:
			walkStmts(info, s.Body.List, echoAlias, filePath, fset, routers, out, hmap)
		case *ast.SelectStmt:
			walkStmts(info, s.Body.List, echoAlias, filePath, fset, routers, out, hmap)
		case *ast.CaseClause:
			walkStmts(info, s.Body, echoAlias, filePath, fset, routers, out, hmap)
		case *ast.CommClause:
			walkStmts(info, s.Body, echoAlias, filePath, fset, routers, out, hmap)
		}
	}
}
