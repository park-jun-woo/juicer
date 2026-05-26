//ff:func feature=scan type=test control=sequence
//ff:what TestResolveResponseType_IdentCov 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveResponseType_IdentCov(t *testing.T) {
	ident := &ast.Ident{Name: "x"}
	info := &types.Info{
		Uses:  make(map[*ast.Ident]types.Object),
		Defs:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	resolveResponseType(ident, info)
}
