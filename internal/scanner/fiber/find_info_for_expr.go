//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 주어진 AST 표현이 속한 패키지의 TypesInfo를 찾는다
package fiber

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/packages"
)

func findInfoForExpr(expr ast.Expr, pkgs []*packages.Package) *types.Info {
	pos := expr.Pos()
	for _, pkg := range pkgs {
		if pkg.TypesInfo == nil {
			continue
		}
		for _, file := range pkg.Syntax {
			if file.Pos() <= pos && pos <= file.End() {
				return pkg.TypesInfo
			}
		}
	}
	return nil
}
