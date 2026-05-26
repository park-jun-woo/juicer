//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_NilTypeCov 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveExprType_NilTypeCov(t *testing.T) {
	ident := &ast.Ident{Name: "x"}
	info := &types.Info{
		Uses:  make(map[*ast.Ident]types.Object),
		Defs:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	tn, _ := resolveExprType(ident, info)
	if tn != "" {
		t.Fatal("expected empty")
	}
}
