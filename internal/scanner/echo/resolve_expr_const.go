//ff:func feature=scan type=extract control=selection
//ff:what 식별자/셀렉터 표현식이 상수이면 그 문자열 값을 반환한다
package echo

import (
	"go/ast"
	"go/types"
)

// resolveExprConst resolves an Ident or SelectorExpr to its const string value.
// The const value is unquoted (ConstToString returns the Go-quoted literal for strings).
func resolveExprConst(info *types.Info, expr ast.Expr) string {
	if info == nil {
		return ""
	}
	var raw string
	switch e := expr.(type) {
	case *ast.Ident:
		raw = resolveUsesConst(info, e)
	case *ast.SelectorExpr:
		raw = resolveUsesConst(info, e.Sel)
	}
	if raw == "" {
		return ""
	}
	return unquote(raw)
}
