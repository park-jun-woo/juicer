//ff:func feature=scan type=test control=sequence
//ff:what TestWalkForGroupArgs_Empty 테스트
package gogin

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestWalkForGroupArgs_Empty(t *testing.T) {
	ctx := &groupArgCtx{
		ginAlias:  "gin",
		routers:   map[string]*routerInfo{},
		fset:      token.NewFileSet(),
		endpoints: []scanner.Endpoint{},
		hmap:      map[int][]ast.Expr{},
		epIndex: map[struct {
			file string
			line int
		}]int{},
	}

	walkForGroupArgs(nil, ctx)
}
