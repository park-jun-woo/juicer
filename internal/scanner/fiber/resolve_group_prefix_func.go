//ff:func feature=scan type=extract control=sequence
//ff:what 함수 하나에서 routers 맵을 구축하고 인라인 Group 인자를 추적한다
package fiber

import (
	"go/ast"

	"golang.org/x/tools/go/packages"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveGroupPrefixFunc(fn *ast.FuncDecl, fiberAlias string, pkg *packages.Package, pkgs []*packages.Package, root string, idx *funcIndex, endpoints []scanner.Endpoint, hmap map[int][]ast.Expr, epIndex map[struct{ file string; line int }]int) {
	routers := make(map[string]*routerInfo)
	registerParams(fn, fiberAlias, routers)
	ctx := &groupArgCtx{
		fiberAlias: fiberAlias,
		routers:    routers,
		info:       pkg.TypesInfo,
		fset:       pkg.Fset,
		idx:        idx,
		root:       root,
		pkgs:       pkgs,
		endpoints:  endpoints,
		hmap:       hmap,
		epIndex:    epIndex,
	}
	walkForGroupArgs(fn.Body.List, ctx)
}
