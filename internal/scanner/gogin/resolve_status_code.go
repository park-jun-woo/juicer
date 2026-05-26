//ff:func feature=scan type=extract control=selection
//ff:what 상태 코드 인자를 정수 리터럴 또는 net/http 상수에서 해석한다
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func resolveStatusCode(expr ast.Expr, info *types.Info) string {
	if lit, ok := expr.(*ast.BasicLit); ok && lit.Kind == token.INT {
		return lit.Value
	}

	if info == nil {
		return "(unknown)"
	}

	switch e := expr.(type) {
	case *ast.SelectorExpr:
		if s := resolveUsesConst(info, e.Sel); s != "" {
			return s
		}
	case *ast.Ident:
		if s := resolveUsesConst(info, e); s != "" {
			return s
		}
	}

	if tv, ok := info.Types[expr]; ok && tv.Value != nil {
		return scanner.ConstToString(tv.Value)
	}

	return "(unknown)"
}
