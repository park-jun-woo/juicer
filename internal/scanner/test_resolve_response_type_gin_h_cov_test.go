//ff:func feature=scan type=test control=sequence
//ff:what TestResolveResponseType_GinHCov 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveResponseType_GinHCov(t *testing.T) {
	comp := &ast.CompositeLit{
		Type: &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "H"}},
	}
	info := &types.Info{
		Uses:  make(map[*ast.Ident]types.Object),
		Defs:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	tn, _, conf := resolveResponseType(comp, info)
	if tn != "gin.H" {
		t.Fatalf("expected gin.H, got %s", tn)
	}
	if conf != "partial" {
		t.Fatalf("expected partial, got %s", conf)
	}
}
