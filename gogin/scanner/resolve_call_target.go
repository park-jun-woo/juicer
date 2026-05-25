//ff:func feature=scan type=extract control=selection
//ff:what 호출 대상 함수의 선언 위치를 해석한다
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
)

// resolveCallTarget resolves the target function position from a call expression.
func resolveCallTarget(call *ast.CallExpr, info *types.Info) token.Pos {
	switch fn := call.Fun.(type) {
	case *ast.SelectorExpr:
		if sel, ok := info.Selections[fn]; ok {
			return sel.Obj().Pos()
		}
		if use, ok := info.Uses[fn.Sel]; ok {
			return use.Pos()
		}
	case *ast.Ident:
		if use, ok := info.Uses[fn]; ok {
			return use.Pos()
		}
	}
	return token.NoPos
}
