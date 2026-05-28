//ff:type feature=scan type=model
//ff:what groupArgCtx 데이터 구조
package echo

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/packages"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

// groupArgCtx holds context for resolving inline Group() arguments in function calls.
type groupArgCtx struct {
	echoAlias string
	routers   map[string]*routerInfo
	info      *types.Info
	fset      *token.FileSet
	idx       *funcIndex
	root      string
	pkgs      []*packages.Package
	endpoints []scanner.Endpoint
	hmap      map[int][]ast.Expr
	epIndex   map[struct {
		file string
		line int
	}]int
}
