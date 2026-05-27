//ff:func feature=scan type=extract control=selection
//ff:what AST 표현의 타입을 go/types로 추적하여 TypeName과 Fields를 반환한다
package gogin

import (
	"go/ast"
	"go/types"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveExprType(expr ast.Expr, info *types.Info) (typeName string, fields []scanner.Field) {
	if info == nil {
		return "", nil
	}

	var t types.Type

	switch e := expr.(type) {
	case *ast.Ident:
		if obj := info.Uses[e]; obj != nil {
			t = obj.Type()
		} else if obj := info.Defs[e]; obj != nil {
			t = obj.Type()
		}
	case *ast.SelectorExpr:
		if obj := info.Uses[e.Sel]; obj != nil {
			t = obj.Type()
		}
	case *ast.CompositeLit:
		if tv, ok := info.Types[e]; ok {
			t = tv.Type
		}
	default:
		if tv, ok := info.Types[expr]; ok {
			t = tv.Type
		}
	}

	if t == nil {
		return "", nil
	}

	return resolveType(t)
}

