//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 파일 내 함수 선언을 순회하며 인라인 Group 인자를 처리한다
package gogin

import (
	"go/ast"

	"golang.org/x/tools/go/packages"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func resolveGroupPrefixFile(file *ast.File, pkg *packages.Package, pkgs []*packages.Package, root string, idx *funcIndex, endpoints []scanner.Endpoint, hmap map[int][]ast.Expr, epIndex map[struct{ file string; line int }]int) {
	ginAlias := ginPkgName(file)
	if ginAlias == "" {
		return
	}
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}
		resolveGroupPrefixFunc(fn, ginAlias, pkg, pkgs, root, idx, endpoints, hmap, epIndex)
	}
}
