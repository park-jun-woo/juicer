//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallTarget_SelectorWithUsesCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallTarget_SelectorWithUsesCov(t *testing.T) {
	selIdent := &ast.Ident{Name: "Func"}
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "pkg"}, Sel: selIdent}
	call := &ast.CallExpr{Fun: sel}
	pkg := types.NewPackage("example.com/pkg", "pkg")
	obj := types.NewFunc(token.Pos(42), pkg, "Func", types.NewSignatureType(nil, nil, nil, nil, nil, false))
	info := &types.Info{
		Uses:       map[*ast.Ident]types.Object{selIdent: obj},
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	pos := resolveCallTarget(call, info)
	if pos != token.Pos(42) {
		t.Fatalf("expected 42, got %d", pos)
	}
}
