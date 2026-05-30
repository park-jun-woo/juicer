//ff:func feature=scan type=test control=sequence topic=echo
//ff:what emptyGroupCtx 테스트 헬퍼
package echo

import (
	"go/ast"
	"go/token"
)

func emptyGroupCtx() *groupArgCtx {
	return &groupArgCtx{
		echoAlias: "echo",
		routers:   map[string]*routerInfo{},
		info:      nil,
		fset:      token.NewFileSet(),
		idx:       buildFuncIndex(nil),
		root:      "/root",
		pkgs:      nil,
		hmap:      map[int][]ast.Expr{},
	}
}
