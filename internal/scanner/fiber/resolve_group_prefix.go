//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 인라인 Group() 인자로 전달된 prefix를 대상 함수의 라우트에 전파한다
package fiber

import (
	"go/ast"

	"golang.org/x/tools/go/packages"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func resolveGroupPrefix(pkgs []*packages.Package, root string, idx *funcIndex, endpoints []scanner.Endpoint, hmap map[int][]ast.Expr) {
	ctxEpIndex := buildEndpointIndex(endpoints)

	for _, pkg := range pkgs {
		if pkg.TypesInfo == nil {
			continue
		}
		for i, file := range pkg.Syntax {
			if i >= len(pkg.CompiledGoFiles) {
				continue
			}
			resolveGroupPrefixFile(file, pkg, pkgs, root, idx, endpoints, hmap, ctxEpIndex)
		}
	}
}
