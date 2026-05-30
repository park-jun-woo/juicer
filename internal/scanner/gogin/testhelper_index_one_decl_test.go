//ff:func feature=scan type=test control=sequence
//ff:what indexOneDecl 테스트 헬퍼: 단일 선언을 색인하고 registerRoutes 호출을 탐색
package gogin

import (
	"go/ast"
	"go/types"
)

// indexOneDecl registers d (if a FuncDecl) into idx and returns the first
// registerRoutes call found, preferring an already-found prev.
func indexOneDecl(idx *funcIndex, info *types.Info, d ast.Decl, prev *ast.CallExpr) *ast.CallExpr {
	fn, ok := d.(*ast.FuncDecl)
	if !ok {
		return prev
	}
	idx.byPos[fn.Name.Pos()] = fn
	idx.info[fn.Name.Pos()] = info
	call := prev
	ast.Inspect(fn, func(n ast.Node) bool {
		c, ok := n.(*ast.CallExpr)
		if !ok || call != nil {
			return true
		}
		if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "registerRoutes" {
			call = c
		}
		return true
	})
	return call
}
