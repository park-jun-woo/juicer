package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestCheckOneDepthCall_NilInfoOld(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "f"}}
	idx := &funcIndex{byPos: make(map[token.Pos]*ast.FuncDecl), info: make(map[token.Pos]*types.Info)}
	checkOneDepthCall(ep, call, "c", nil, idx)
}
