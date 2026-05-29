//ff:func feature=scan type=extract control=selection
//ff:what collectStringParts 함수
package echo

import (
	"go/ast"
	"go/token"
	"go/types"
)

func collectStringParts(info *types.Info, expr ast.Expr, parts *[]string) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			*parts = append(*parts, unquote(e.Value))
		}
	case *ast.BinaryExpr:
		if e.Op == token.ADD {
			collectStringParts(info, e.X, parts)
			collectStringParts(info, e.Y, parts)
		}
	case *ast.SelectorExpr, *ast.Ident:
		// const reference like config.APIBooks — resolve via types.Info
		if v := resolveExprConst(info, e); v != "" {
			*parts = append(*parts, v)
		}
	}
}
