//ff:func feature=scan type=extract control=selection
//ff:what collectStringParts 함수
package scanner

import (
	"go/ast"
	"go/token"
)

func collectStringParts(expr ast.Expr, parts *[]string) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			*parts = append(*parts, unquote(e.Value))
		}
	case *ast.BinaryExpr:
		if e.Op == token.ADD {
			collectStringParts(e.X, parts)
			collectStringParts(e.Y, parts)
		}
	}
}
